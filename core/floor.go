package core

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/m6yf/bcwork/bcdb"
	"github.com/m6yf/bcwork/bcdb/filter"
	"github.com/m6yf/bcwork/bcdb/order"
	"github.com/m6yf/bcwork/bcdb/pagination"
	"github.com/m6yf/bcwork/bcdb/qmods"
	"github.com/m6yf/bcwork/models"
	"github.com/m6yf/bcwork/utils"
	"github.com/m6yf/bcwork/utils/bcguid"
	"github.com/rotisserie/eris"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

type FloorUpdateRequest struct {
	Publisher string  `json:"publisher"`
	Domain    string  `json:"domain"`
	Device    string  `json:"device"`
	Floor     float64 `json:"floor"`
	Country   string  `json:"country"`
}

type Floor struct {
	Publisher string  `boil:"publisher" json:"publisher" toml:"publisher" yaml:"publisher"`
	Domain    string  `boil:"domain" json:"domain" toml:"domain" yaml:"domain"`
	Country   string  `boil:"country" json:"country" toml:"country" yaml:"country"`
	Device    string  `boil:"device" json:"device" toml:"device" yaml:"device"`
	Floor     float64 `boil:"floor" json:"floor" toml:"floor" yaml:"floor"`
}

type FloorSlice []*Floor

type GetFloorOptions struct {
	Filter     FloorFilter            `json:"filter"`
	Pagination *pagination.Pagination `json:"pagination"`
	Order      order.Sort             `json:"order"`
	Selector   string                 `json:"selector"`
}

type FloorFilter struct {
	Publisher filter.StringArrayFilter `json:"publisher"`
	Domain    filter.StringArrayFilter `json:"domain"`
	Country   filter.StringArrayFilter `json:"country"`
	Device    filter.StringArrayFilter `json:"device"`
}

type FloorRealtimeRecord struct {
	Rule    string  `json:"rule"`
	Floor   float64 `json:"floor"`
	FloorID string  `json:"floor_id"`
}

func (f FloorUpdateRequest) GetPublisher() string { return f.Publisher }
func (f FloorUpdateRequest) GetDomain() string    { return f.Domain }
func (f FloorUpdateRequest) GetDevice() string    { return f.Device }
func (f FloorUpdateRequest) GetCountry() string   { return f.Country }

func (floor *Floor) FromModel(mod *models.Floor) error {

	floor.Publisher = mod.Publisher
	floor.Domain = mod.Domain
	floor.Country = mod.Country
	floor.Device = mod.Device
	floor.Floor = mod.Floor

	return nil
}

func (cs *FloorSlice) FromModel(slice models.FloorSlice) error {

	for _, mod := range slice {
		c := Floor{}
		err := c.FromModel(mod)
		if err != nil {
			return eris.Cause(err)
		}
		*cs = append(*cs, &c)
	}

	return nil
}

func GetFloors(ctx context.Context, ops *GetFloorOptions) (FloorSlice, error) {

	qmods := ops.Filter.QueryMod().Order(ops.Order, nil, models.FloorColumns.Publisher).AddArray(ops.Pagination.Do())

	qmods = qmods.Add(qm.Select("DISTINCT *"))

	mods, err := models.Floors(qmods...).All(ctx, bcdb.DB())
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, eris.Wrap(err, "failed to retrieve floors")
	}

	res := make(FloorSlice, 0)
	err = res.FromModel(mods)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (filter *FloorFilter) QueryMod() qmods.QueryModsSlice {

	mods := make(qmods.QueryModsSlice, 0)

	if filter == nil {
		return mods
	}

	if len(filter.Publisher) > 0 {
		mods = append(mods, filter.Publisher.AndIn(models.FloorColumns.Publisher))
	}

	if len(filter.Device) > 0 {
		mods = append(mods, filter.Device.AndIn(models.FloorColumns.Device))
	}

	if len(filter.Domain) > 0 {
		mods = append(mods, filter.Domain.AndIn(models.FloorColumns.Domain))
	}

	if len(filter.Country) > 0 {
		mods = append(mods, filter.Country.AndIn(models.FloorColumns.Country))
	}

	return mods
}

func UpdateFloorMetaData(c *fiber.Ctx, data *FloorUpdateRequest) error {
	_, err := json.Marshal(data)

	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to parse hash value for floor")
	}

	err = SendFloorToRT(context.Background(), *data)
	if err != nil {
		return err
	}
	return nil
}

func UpdateFloors(c *fiber.Ctx, data *FloorUpdateRequest) error {

	modConf := models.Floor{
		Publisher: data.Publisher,
		Domain:    data.Domain,
		Device:    data.Device,
		Floor:     data.Floor,
		Country:   data.Country,
	}

	return modConf.Upsert(c.Context(), bcdb.DB(), true, []string{models.FloorColumns.Publisher, models.FloorColumns.Domain, models.FloorColumns.Device, models.FloorColumns.Country}, boil.Infer(), boil.Infer())
}

func SendFloorToRT(c context.Context, updateRequest FloorUpdateRequest) error {

	modFloor, err := floorQuery(c, updateRequest)

	if err != nil && err != sql.ErrNoRows {
		return eris.Wrapf(err, "failed to fetch floors")
	}

	var finalRules []FloorRealtimeRecord

	finalRules = createFloorMetadata(modFloor, finalRules, updateRequest)

	finalOutput := struct {
		Rules []FloorRealtimeRecord `json:"rules"`
	}{Rules: finalRules}

	value, err := json.Marshal(finalOutput)
	if err != nil {
		return eris.Wrap(err, "failed to marshal floorRT to JSON")
	}

	key := utils.GetMetadataKey(updateRequest)
	metadataKey := utils.CreateMetadataKey(key, "price:floor:v2")
	metadataValue := CreateMetadataValueFloor(updateRequest, metadataKey, value)

	err = metadataValue.Insert(c, bcdb.DB(), boil.Infer())
	if err != nil {
		return eris.Wrap(err, "failed to insert metadata record for floor")
	}

	return nil
}

func floorQuery(c context.Context, updateRequest FloorUpdateRequest) (models.FloorSlice, error) {
	modFloor, err := models.Floors(
		models.FloorWhere.Country.EQ(updateRequest.Country),
		models.FloorWhere.Domain.EQ(updateRequest.Domain),
		models.FloorWhere.Device.EQ(updateRequest.Device),
		models.FloorWhere.Publisher.EQ(updateRequest.Publisher),
	).All(c, bcdb.DB())
	return modFloor, err
}

func createFloorMetadata(modFloor models.FloorSlice, finalRules []FloorRealtimeRecord, updateRequest FloorUpdateRequest) []FloorRealtimeRecord {
	if len(modFloor) != 0 {
		floors := make(FloorSlice, 0)
		floors.FromModel(modFloor)

		for _, floor := range floors {
			rule := FloorRealtimeRecord{
				Rule:    utils.GetFormulaRegex(floor.Country, floor.Domain, floor.Device, false),
				Floor:   floor.Floor,
				FloorID: floor.Publisher,
			}
			finalRules = append(finalRules, rule)
		}
	}

	newRule := FloorRealtimeRecord{
		Rule:    utils.GetFormulaRegex(updateRequest.Country, updateRequest.Domain, updateRequest.Device, false),
		Floor:   updateRequest.Floor,
		FloorID: updateRequest.Publisher,
	}
	finalRules = append(finalRules, newRule)
	return finalRules
}

//func getMetadataKeyFloor(updateRequest FloorUpdateRequest) utils.MetadataKey {
//	key := utils.MetadataKey{
//		Publisher: updateRequest.Publisher,
//		Domain:    updateRequest.Domain,
//		Device:    updateRequest.Device,
//		Country:   updateRequest.Country,
//	}
//	return key
//}

func CreateMetadataValueFloor(updateRequest FloorUpdateRequest, key string, b []byte) models.MetadataQueue {
	modMeta := models.MetadataQueue{
		TransactionID: bcguid.NewFromf(updateRequest.Publisher, updateRequest.Domain, time.Now()),
		Key:           key,
		Value:         b,
	}
	return modMeta
}
