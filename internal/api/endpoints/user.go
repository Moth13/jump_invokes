package endpoints

import (
	"fmt"
	handlers "invokes/internal/api/handlers"
	"net/http"

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
		users, count, err := e.DB.GetUsers()

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
