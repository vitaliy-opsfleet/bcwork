package core

import (
	"context"
	"database/sql"
	"encoding/json"
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

type ConfiantUpdateRequest struct {
	Publisher string  `json:"publisher_id" validate:"required"`
	Domain    string  `json:"domain"`
	Hash      string  `json:"confiant_key"`
	Rate      float64 `json:"rate"`
}

type ConfiantUpdateRespose struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Confiant struct {
	ConfiantKey string     `boil:"confiant_key" json:"confiant_key" toml:"confiant_key" yaml:"confiant_key"`
	PublisherID string     `boil:"publisher_id" json:"publisher_id" toml:"publisher_id" yaml:"publisher_id"`
	Domain      string     `boil:"domain" json:"domain,omitempty" toml:"domain" yaml:"domain,omitempty"`
	Rate        float64    `boil:"rate" json:"rate" toml:"rate" yaml:"rate"`
	CreatedAt   time.Time  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   *time.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
}

type ConfiantSlice []*Confiant

type GetConfiantOptions struct {
	Filter     ConfiantFilter         `json:"filter"`
	Pagination *pagination.Pagination `json:"pagination"`
	Order      order.Sort             `json:"order"`
	Selector   string                 `json:"selector"`
}

type ConfiantFilter struct {
	PublisherID filter.StringArrayFilter `json:"publisher_id,omitempty"`
	ConfiantID  filter.StringArrayFilter `json:"confiant_key,omitempty"`
	Domain      filter.StringArrayFilter `json:"domain,omitempty"`
	Rate        filter.StringArrayFilter `json:"rate,omitempty"`
}

func (confiant *Confiant) FromModel(mod *models.Confiant) error {

	confiant.PublisherID = mod.PublisherID
	confiant.CreatedAt = mod.CreatedAt
	confiant.Domain = mod.Domain
	confiant.Rate = mod.Rate
	confiant.ConfiantKey = mod.ConfiantKey

	return nil
}

func (cs *ConfiantSlice) FromModel(slice models.ConfiantSlice) error {

	for _, mod := range slice {
		c := Confiant{}
		err := c.FromModel(mod)
		if err != nil {
			return eris.Cause(err)
		}
		*cs = append(*cs, &c)
	}

	return nil
}

func GetConfiants(ctx context.Context, ops *GetConfiantOptions) (ConfiantSlice, error) {

	qmods := ops.Filter.QueryMod().Order(ops.Order, nil, models.ConfiantColumns.PublisherID).AddArray(ops.Pagination.Do())

	if ops.Selector == "id" {
		qmods = qmods.Add(qm.Select("DISTINCT " + models.ConfiantColumns.PublisherID))
	} else {
		qmods = qmods.Add(qm.Select("DISTINCT *"))
		qmods = qmods.Add(qm.Load(models.ConfiantRels.Publisher))

	}
	mods, err := models.Confiants(qmods...).All(ctx, bcdb.DB())
	if err != nil && err != sql.ErrNoRows {
		return nil, eris.Wrap(err, "failed to retrieve publishers")
	}

	res := make(ConfiantSlice, 0)
	res.FromModel(mods)

	return res, nil
}

func (filter *ConfiantFilter) QueryMod() qmods.QueryModsSlice {

	mods := make(qmods.QueryModsSlice, 0)

	if filter == nil {
		return mods
	}

	if len(filter.PublisherID) > 0 {
		mods = append(mods, filter.PublisherID.AndIn(models.ConfiantColumns.PublisherID))
	}

	if len(filter.ConfiantID) > 0 {
		mods = append(mods, filter.ConfiantID.AndIn(models.ConfiantColumns.ConfiantKey))
	}

	if len(filter.Domain) > 0 {
		mods = append(mods, filter.Domain.AndIn(models.ConfiantColumns.Domain))
	}

	if len(filter.Rate) > 0 {
		mods = append(mods, filter.Rate.AndIn(models.ConfiantColumns.Rate))
	}

	return mods
}

func UpdateMetaDataQueue(c *fiber.Ctx, data *ConfiantUpdateRequest) error {

	val, err := json.Marshal(data.Hash)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Confiant failed to parse hash value")
	}

	mod := models.MetadataQueue{
		Key:           "confiant:" + data.Publisher,
		TransactionID: bcguid.NewFromf(data.Publisher, data.Domain, time.Now()),
		Value:         val,
	}

	if data.Domain != "" {
		mod.Key = mod.Key + ":" + data.Domain
	}

	err = mod.Insert(c.Context(), bcdb.DB(), boil.Infer())

	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to insert metadata update to queue")
	}
	return nil
}

func UpdateConfiant(c *fiber.Ctx, data *ConfiantUpdateRequest) error {

	modConf := models.Confiant{
		PublisherID: data.Publisher,
		ConfiantKey: data.Hash,
		Rate:        data.Rate,
		Domain:      data.Domain,
	}

	return modConf.Upsert(c.Context(), bcdb.DB(), true, []string{models.ConfiantColumns.PublisherID, models.ConfiantColumns.Domain}, boil.Infer(), boil.Infer())
}
