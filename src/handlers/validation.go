package handlers

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/Eli15x/search-car/src/service"
	"github.com/gin-gonic/gin"
)

func ValidateNumber(c *gin.Context) {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	userId := json_map["userId"].(string)
	number := json_map["number"].(string)

	if userId == "" {
		c.String(http.StatusBadRequest, "Validate User Error: userId not find")
		return
	}

	if number == "" {
		c.String(http.StatusBadRequest, "Validate User Error: number not find")
		return
	}

	err = service.GetInstanceValidation().ValidateNumber(userId, number)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resultUser, err := service.GetInstanceUser().GetInformation(context.Background(), userId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, resultUser)
}
