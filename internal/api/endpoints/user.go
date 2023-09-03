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
// @Tags user
// @ID get-users
// @Summary Get a list of users
// @Description Retreive a list of users
// @Produce application/json
// @Param user_id query int false "user id"
// @Param balance query string false "user balance"
// @Param first_name query string false "user first name"
// @Param last_name query string false "user last name"
// @Success 200 {object} models.Users
// @Failure 204 {string} string "No content found"
// @Failure 400 {string} string "Invalid request parameters"
// @Failure 500 {object} string "Internal server error"
// @Router /users [get]
func GetUsers(e *handlers.Env) gin.HandlerFunc {

	return func(c *gin.Context) {
		filter := models.User{}
		userID := c.DefaultQuery("user_id", "")
		balance := c.DefaultQuery("balance", "")
		firstName := c.DefaultQuery("first_name", "")
		lastName := c.DefaultQuery("last_name", "")
		if userID != "" {
			if uid, err := strconv.Atoi(userID); err == nil {
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
		if firstName != "" {
			filter.FirstName = firstName
		}
		if lastName != "" {
			filter.LastName = lastName
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
