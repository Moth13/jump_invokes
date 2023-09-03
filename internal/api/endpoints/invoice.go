package endpoints

import (
	"fmt"
	handlers "invokes/internal/api/handlers"
	"invokes/internal/db"
	"invokes/internal/models"
	"invokes/internal/utils"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PostInvoiceResponse struct {
	Msg       string `json:"invoice_id"`
	InvoiceID int    `json:"invoice_id"`
}

// PostInvoice godoc
// @Tags invoice
// @ID post-invoice
// @Summary Post an invoice
// @Description Post an invoice
// @Produce application/json
// @Success 200 {string} string "pong"
// @Failure 500 {object} string "Internal server error"
// @Router /invoice [post]
func PostInvoice(e *handlers.Env) gin.HandlerFunc {

	return func(c *gin.Context) {

		invoice := models.Invoice{}

		if err := c.ShouldBindBodyWith(&invoice, binding.JSON); err != nil {
			utils.Logger.Info(err)
			c.JSON(http.StatusBadRequest, fmt.Sprintf("err %s", err))
			return
		}

		if invoice.AmountFloat != math.Floor(invoice.AmountFloat*100)/100 {
			c.JSON(http.StatusBadRequest, "Incorrect Amount floor value, has to be in 2 decimals")
			return
		}

		httpcode := http.StatusNoContent
		msg := ""
		err := e.DB.AddInvoice(&invoice)
		if err != nil {
			httpcode = http.StatusInternalServerError
			switch e := err.(type) {
			case *db.DBError:
				httpcode = utils.DBCodeToHTTPCode(err.(*db.DBError).Type)
			default:
				log.Println(e)
			}
			msg = fmt.Sprintf("Error %s", err)
		}
		c.JSON(httpcode, msg)
	}
}

// GetInvoices godoc
// @Tags Invoices
// @ID get-Invoices
// @Summary Get a list of Invoices
// @Description Retreive a list of Invoices
// @Produce application/json
// @Success 200 {object} models.Invoices
// @Failure 204 {object} string "No content found"
// @Failure 500 {object} string "Internal server error"
// @Router /Invoices [get]
// @x-codeSamples [{"lang":"Shell","label":"cURL","source":"curl --include \\\n     --header \"Content-type: application/json\" \\\n     -X GET \"{server_url}/Invoices\"\n"},{"lang":"Python","source":"import requests\nh = {\n  \"Content-type\": \"application/json\"\n}\np = {}\nresp = requests.get(\"{server_url}/Invoices\", params=p, headers=h)\n"},{"lang":"JavaScript","source":"var request = require('request');\nrequest({\n  method: 'GET',\n  url: '{server_url}/Invoices',\n  headers: {\n    'Content-Type': 'application/json',\n  }}, function (error, response, body) {\n  console.log('Status:', response.statusCode);\n  console.log('Headers:', JSON.stringify(response.headers));\n  console.log('Response:', body);\n});\n"}]
func GetInvoices(e *handlers.Env) gin.HandlerFunc {

	return func(c *gin.Context) {
		filter := models.Invoice{}
		invoice_id := c.DefaultQuery("invoice_id", "")
		user_id := c.DefaultQuery("user_id", "")
		amount := c.DefaultQuery("amount", "")
		label := c.DefaultQuery("label", "")
		if invoice_id != "" {
			if iid, err := strconv.Atoi(invoice_id); err == nil {
				filter.ID = iid
			} else {
				c.JSON(http.StatusBadRequest, "Invalid invoice_id format")
				return
			}
		}
		if user_id != "" {
			if uid, err := strconv.Atoi(user_id); err == nil {
				filter.UserID = uid
			} else {
				c.JSON(http.StatusBadRequest, "Invalid user_id format")
				return
			}
		}
		if amount != "" {
			if a, err := strconv.ParseFloat(amount, 64); err == nil {
				filter.Amount = int32(a * 100)
			} else {
				c.JSON(http.StatusBadRequest, "Invalid amount format")
				return
			}
		}
		if label != "" {
			filter.Label = label
		}

		invoices, count, err := e.DB.GetInvoices(&filter)

		// If encounter an error, failed it
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("Internal server error %s", err))
			// If no count found
		} else if count == 0 {
			c.JSON(http.StatusNoContent, "No invoices found")
		} else {
			c.JSON(http.StatusOK, invoices)
		}
	}
}
