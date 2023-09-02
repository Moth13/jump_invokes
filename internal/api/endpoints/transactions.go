package endpoints

import (
	"fmt"
	handlers "invokes/internal/api/handlers"
	"invokes/internal/db"
	"invokes/internal/models"
	"invokes/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// PostTransaction godoc
// @Tags Transaction
// @ID post-Transaction
// @Summary Post an Transaction
// @Description Post an Transaction
// @Produce application/json
// @Success 200 {string} string "pong"
// @Failure 500 {object} string "Internal server error"
// @Router /Transaction [post]
func PostTransaction(e *handlers.Env) gin.HandlerFunc {

	return func(c *gin.Context) {

		Transaction := models.Transaction{}

		if err := c.ShouldBindBodyWith(&Transaction, binding.JSON); err != nil {
			utils.Logger.Info(err)
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error %s", err))
			return
		}

		httpcode := http.StatusNoContent
		msg := ""
		err := e.DB.AddTransaction(&Transaction)
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
