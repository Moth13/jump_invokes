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

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// PostTransaction godoc
// @Tags transaction
// @ID post-transaction
// @Summary Post an Transaction
// @Description Post an Transaction
// @Produce application/json
// @Success 204 {string} string ""
// @Failure 400 {string} string "Invalid post parameters"
// @Failure 404 {string} string "No invoice found"
// @Failure 409 {string} string "Invoice already exist"
// @Failure 500 {object} string "Internal server error"
// @Router /Transaction [post]
func PostTransaction(e *handlers.Env) gin.HandlerFunc {

	return func(c *gin.Context) {

		transaction := models.Transaction{}

		if err := c.ShouldBindBodyWith(&transaction, binding.JSON); err != nil {
			utils.Logger.Info(err)
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error %s", err))
			return
		}

		if transaction.AmountFloat != math.Floor(transaction.AmountFloat*100)/100 {
			c.JSON(http.StatusBadRequest, "Incorrect Amount floor value, has to be in 2 decimals")
			return
		}

		httpcode := http.StatusNoContent
		msg := ""
		err := e.DB.AddTransaction(&transaction)
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
