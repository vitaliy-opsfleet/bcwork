package nbsupply

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"github.com/friendsofgo/errors"
	"github.com/m6yf/bcwork/models"
	"github.com/m6yf/bcwork/utils/bcguid"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type CompassNewBidderRecord struct {
	CombinationId          string  `json:"CombinationId"`
	DateStamp              int64   `json:"DateStamp"`
	PublisherId            int64   `json:"PublisherId"`
	Domain                 string  `json:"Domain"`
	Country                string  `json:"Country"`
	Device                 string  `json:"Device"`
	Os                     string  `json:"Os"`
	Size                   string  `json:"Size,omitempty"`
	PaymentType            string  `json:"PaymentType,omitempty"`
	PublisherTagName       string  `json:"PublisherTagName,omitempty"`
	PublisherPlacementType string  `json:"PublisherPlacementType"`
	PublisherBidRequests   int64   `json:"PublisherBidRequests"`
	BCMBidResponses        int64   `json:"BCMBidResponses"`
	SoldImps               int64   `json:"SoldImps"`
	PubImps                int64   `json:"PubImps"`
	Cost                   float64 `json:"Cost"`
	Revenue                float64 `json:"Revenue"`
	AvgBidPrice            float64 `json:"AvgBidPrice"`
	LoopingRatio           float64 `json:"LoopingRatio"`
	UpdatedAt              string  `json:"updatedAt"`
	DataImpressions        int64   `json:"DataImpressions"`
	DataFee                float64 `json:"DataFee"`
}

var loc *time.Location

func ConvertToCompass(modSlice models.NBSupplyHourlySlice) []*CompassNewBidderRecord {
	var err error
	if loc == nil {
		loc, err = time.LoadLocation("EST")
		if err != nil {
			log.Fatal().Err(err).Msg("failed to find EST loc")
		}
	}

	var res []*CompassNewBidderRecord
	for _, mod := range modSlice {

		val := &CompassNewBidderRecord{
			Domain:                 mod.Domain,
			Country:                mod.Country,
			Os:                     mod.Os,
			PublisherPlacementType: mod.PlacementType,
			Device:                 mod.DeviceType,
			PublisherBidRequests:   mod.BidRequests,
			BCMBidResponses:        mod.BidResponses,
			Cost:                   mod.Cost,
			Revenue:                mod.Revenue - mod.DemandPartnerFee,
			PubImps:                mod.PublisherImpressions,
			SoldImps:               mod.SoldImpressions,
			AvgBidPrice:            mod.AvgBidPrice,
			UpdatedAt:              time.Now().In(loc).Format("2006-01-02 15:04:00"),
			LoopingRatio:           float64(mod.SoldImpressions + mod.MissedOpportunities),
			DataImpressions:        mod.DataImpressions,
			DataFee:                mod.DataFee,
		}

		//if true {
		val.Size = mod.Size
		if mod.RequestType != "js" {
			val.Size = "-"
		}

		val.PaymentType = "NP HB"
		if mod.RequestType == "js" {
			val.PaymentType = "NP CPM"
		}

		val.PublisherTagName = "Header Bidding"
		if mod.RequestType == "js" {
			val.PublisherTagName = val.Domain + "_" + val.Device + "_" + val.Size
		}

		//}

		val.DateStamp = mod.Time.In(loc).Unix() / 100
		val.CombinationId = bcguid.NewFromf(mod.PublisherID, mod.Domain, mod.Country, mod.DeviceType, mod.Os, mod.PlacementType, mod.RequestType, mod.Size, mod.PaymentType, val.DateStamp)
		val.PublisherId, err = strconv.ParseInt(mod.PublisherID, 10, 64)
		if err != nil {
			log.Error().Err(err).Interface("mod", mod).Msgf("illegal publisher id when parsing to int, 0 will be used")
		}

		res = append(res, val)
	}

	return res
}

func Send(ctx context.Context, vals []*CompassNewBidderRecord) error {

	b, err := json.Marshal(vals)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal compass values")
	}

	var buf bytes.Buffer
	g := gzip.NewWriter(&buf)
	if err != nil {
		return errors.Wrapf(err, "failed to create gzip writer")
	}

	if _, err := g.Write(b); err != nil {
		return errors.Wrapf(err, "failed to write to gzip writer")
	}
	if err = g.Close(); err != nil {
		return errors.Wrapf(err, "failed to close gzip writer")
	}

	log.Info().Int("payload.records", len(vals)).Int("payload.bytes", len(b)).Int("payload.compressed", buf.Len()).Msg("sending to compass")
	req, err := http.NewRequest("POST", "https://nb-reports.ministerial5.com/supply-reports", &buf)
	//req, err := http.NewRequest("POST", "https://staging-nb-reports.ministerial5.com/supply-reports", &buf)
	req.Header.Add("Content-Encoding", "gzip")
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return errors.Wrapf(err, "failed to send post requests")
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "failed to read post body")
	}
	if resp.StatusCode != http.StatusOK {
		log.Error().Str("rec", string(b)).Msg("post request returned http error")
		return errors.Wrapf(err, "http error")
	}

	//i += pageSize
	//log.Info().Msgf("page %d sent", i/pageSize)
	//
	return nil

}

//func SendStaging(ctx context.Context, vals []*CompassNewBidderRecord) error {
//
//	b, err := json.Marshal(vals)
//	if err != nil {
//		return errors.Wrapf(err, "failed to marshal compass values")
//	}
//
//	var buf bytes.Buffer
//	g := gzip.NewWriter(&buf)
//	if err != nil {
//		return errors.Wrapf(err, "failed to create gzip writer")
//	}
//
//	if _, err := g.Write(b); err != nil {
//		return errors.Wrapf(err, "failed to write to gzip writer")
//	}
//	if err = g.Close(); err != nil {
//		return errors.Wrapf(err, "failed to close gzip writer")
//	}
//
//	log.Info().Int("payload.records", len(vals)).Int("payload.bytes", len(b)).Int("payload.compressed", buf.Len()).Msg("sending to staging")
//	req, err := http.NewRequest("POST", "https://nb-reports.ministerial5.com/staging-reports", &buf)
//	req.Header.Add("Content-Encoding", "gzip")
//	req.Header.Add("Content-Type", "application/json")
//	client := http.Client{}
//	resp, err := client.Do(req)
//	defer resp.Body.Close()
//	if err != nil {
//		return errors.Wrapf(err, "failed to send post requests")
//	}
//
//	b, err = ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return errors.Wrapf(err, "failed to read post body")
//	}
//	if resp.StatusCode != http.StatusOK {
//		log.Error().Str("rec", string(b)).Msg("post request to stageing returned http error")
//		return errors.Wrapf(err, "http error")
//	}
//
//	//i += pageSize
//	//log.Info().Msgf("page %d sent", i/pageSize)
//	//
//	return nil
//
//}
