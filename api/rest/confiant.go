package rest

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/m6yf/bcwork/bcdb"
	"github.com/m6yf/bcwork/models"
	"github.com/m6yf/bcwork/utils/bcguid"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasttemplate"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"golang.org/x/text/message"
	"net/http"
	"time"
)

// DemandReportGetRequest contains filter parameters for retrieving events
type ConfiantUpdateRequest struct {
	Publisher string  `json:"publisher"`
	Domain    string  `json:"domain"`
	Hash      string  `json:"hash"`
	Rate      float64 `json:"rate"`
}

// ConfiantUpdateRespose
type ConfiantUpdateRespose struct {
	// in: body
	Status string `json:"status"`
}

// ConfiantPostHandler Update and enable Confiant setup
// @Description Update and enable Confiant setup (publisher is mandatory, domain is optional)
// @Tags metadata
// @Accept json
// @Produce json
// @Param options body ConfiantUpdateRequest true "Confiant update Options"
// @Success 200 {object} ConfiantUpdateRespose
// @Security ApiKeyAuth
// @Router /confiant [post]
func ConfiantPostHandler(c *fiber.Ctx) error {

	data := &ConfiantUpdateRequest{}
	if err := c.BodyParser(&data); err != nil {
		log.Error().Err(err).Str("body", string(c.Body())).Msg("failed to parse metadata update payload")

		return c.SendStatus(http.StatusBadRequest)
	}

	if data.Publisher == "" {
		c.SendString("'publisher' is mandatory")
		return c.SendStatus(http.StatusBadRequest)
	}

	val, err := json.Marshal(data.Hash)
	if err != nil {
		log.Error().Err(err).Str("body", string(c.Body())).Msg("failed to parse value")
		return c.SendStatus(http.StatusInternalServerError)
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
		log.Error().Err(err).Str("body", string(c.Body())).Msg("failed to insert metadata update to queue")
		return c.SendStatus(http.StatusInternalServerError)
	}

	modConf := models.Confiant{
		PublisherID: data.Publisher,
		ConfiantKey: data.Hash,
		Rate:        1,
	}
	if data.Domain != "" {
		modConf.Domain = null.StringFrom(data.Domain)
	}

	err = modConf.Upsert(c.Context(), bcdb.DB(), true, []string{models.ConfiantColumns.ConfiantKey}, boil.Infer(), boil.Infer())
	if err != nil {
		log.Error().Err(err).Str("body", string(c.Body())).Msg("failed to insert metadata update to queue")
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(ConfiantUpdateRespose{
		Status: "ok",
	})
}

// ConfiantGetHandler Get confiant setup
// @Description Get confiant setup
// @Tags metadata
// @Accept json
// @Produce html
// @Security ApiKeyAuth
// @Router /confiant [get]
func ConfiantGetAllHandler(c *fiber.Ctx) error {

	query := `select metadata_queue.*
from metadata_queue,(select key,max(created_at) created_at from metadata_queue where key like '%confiant%' group by key) last
where last.created_at=metadata_queue.created_at and last.key=metadata_queue.key order by metadata_queue.key`

	records := models.MetadataQueueSlice{}
	err := queries.Raw(query).Bind(c.Context(), bcdb.DB(), &records)
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch all price factors")
		c.SendString("failed to fetch")
		return c.SendStatus(http.StatusInternalServerError)
	}

	if c.Query("format") == "html" {
		c.Set("Content-Type", "text/html")
		b := bytes.Buffer{}
		p := message.NewPrinter(message.MatchLanguage("en"))

		for _, rec := range records {
			b.WriteString(p.Sprintf(rowConfiant, rec.Key, rec.Value, rec.CreatedAt.Format("2006-01-02 15:04"), rec.CommitedInstances))
		}
		t := fasttemplate.New(htmlConfiant, "{{", "}}")
		s := t.ExecuteString(map[string]interface{}{
			"data": b.String(),
		})
		return c.SendString(s)
	} else {
		c.Set("Content-Type", "application/json")
		return c.JSON(records)
	}

}

var htmlConfiant = `
<html>
<head>
     <link href="https://unpkg.com/tailwindcss@^1.0/dist/tailwind.min.css" rel="stylesheet">
</head>
<body>
<div class="md:flex justify-center md:items-center">
   <div class="mt-1 flex md:mt-0 md:ml-4">
    <img class="filter invert h-40 w-40" src="https://onlinemediasolutions.com/wp-content/themes/brightcom/assets/images/oms-logo.svg" alt="">
  </div>
<div class="min-w-0">
    <h2 class="p-3 text-2xl font-bold leading-7 text-purple-600 sm:text-3xl sm:truncate">
      Confiant Setup Status 
    </h2>
  </div>
 
</div>


<div class="flex flex-col">
  <div class="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
    <div class="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
      <div class="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th scope="col" class="font-bold px-6 py-3 text-left text-xs font-medium text-gray-900 uppercase tracking-wider">
                Key
              </th>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-900 uppercase tracking-wider">
                  Hash
               </th>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-900 uppercase tracking-wider">
                  Create At
               </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-900 uppercase tracking-wider">
                  Committed
               </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
              {{data}}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>
</body>
</html>`

var rowConfiant = `<tr>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                     %s
                 </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                     %s
                  </td>
                   <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                     %s
                 </td>
                 <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                     %d
                 </td>
                        
            </tr>`
