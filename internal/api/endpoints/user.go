package endpoints

import (
	"fmt"
	handlers "invokes/internal/api/handlers"
	"invokes/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Tags users
// @ID get-users
// @Summary Get a list of users
// @Description Retreive a list of users
// @Produce application/json
// @Success 200 {object} models.Users
// @Failure 204 {object} string "No content found"
// @Failure 500 {object} string "Internal server error"
// @Router /users [get]
// @x-codeSamples [{"lang":"Shell","label":"cURL","source":"curl --include \\\n     --header \"Content-type: application/json\" \\\n     -X GET \"{server_url}/users\"\n"},{"lang":"Python","source":"import requests\nh = {\n  \"Content-type\": \"application/json\"\n}\np = {}\nresp = requests.get(\"{server_url}/users\", params=p, headers=h)\n"},{"lang":"JavaScript","source":"var request = require('request');\nrequest({\n  method: 'GET',\n  url: '{server_url}/users',\n  headers: {\n    'Content-Type': 'application/json',\n  }}, function (error, response, body) {\n  console.log('Status:', response.statusCode);\n  console.log('Headers:', JSON.stringify(response.headers));\n  console.log('Response:', body);\n});\n"}]
func GetUsers(e *handlers.Env) gin.HandlerFunc {

	return func(c *gin.Context) {
		filter := models.User{}
		user_id := c.DefaultQuery("user_id", "")
		balance := c.DefaultQuery("balance", "")
		first_name := c.DefaultQuery("first_name", "")
		last_name := c.DefaultQuery("last_name", "")
		if user_id != "" {
			if uid, err := strconv.Atoi(user_id); err == nil {
				filter.UserID = uid
			} else {
				c.JSON(http.StatusBadRequest, "Invalid user_id format")
				return
			}
		}
		if balance != "" {
			if a, err := strconv.ParseFloat(balance, 64); err == nil {
				filter.Balance = int32(a * 100)
			} else {
				c.JSON(http.StatusBadRequest, "Invalid balance format")
				return
			}
		}
		if first_name != "" {
			filter.FirstName = first_name
		}
		if last_name != "" {
			filter.LastName = last_name
		}
		users, count, err := e.DB.GetUsers(&filter)

		// If encounter an error, failed it
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("Internal server error %s", err))
			// If no count found
		} else if count == 0 {
			c.JSON(http.StatusNoContent, "No users found")
		} else {
			c.JSON(http.StatusOK, users)
		}
	}
}
