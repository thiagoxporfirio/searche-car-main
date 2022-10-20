package handlers

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/Eli15x/search-car/src/service"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func ValidateUser(c *gin.Context) {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	password := json_map["password"].(string)
	email := json_map["email"].(string)

	if email == "" {
		c.String(http.StatusBadRequest, "Validate User Error: email not find")
		return
	}

	if password == "" {
		c.String(http.StatusBadRequest, "Create User Error: password not find")
		return
	}

	resultUser, err := service.GetInstanceUser().ValidateUser(context.Background(), email, password)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, resultUser)
}

func CreateUser(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	name := json_map["name"].(string)
	email := json_map["email"].(string) //está dando erro quando tenta pegar o "email" e ele não existe.
	password := json_map["password"].(string)
	telefone := json_map["telefone"].(string)
	username := json_map["username"].(string)

	if name == "" {
		c.String(http.StatusBadRequest, "Create User Error: name not find")
		return
	}

	if username == "" {
		c.String(http.StatusBadRequest, "Create User Error: username not find")
		return
	}

	if email == "" {
		c.String(400, "Create User Error: email not find")
		return
	}

	if password == "" {
		c.String(400, "Create User Error: password not find")
		return
	}

	userId, err := service.GetInstanceUser().CreateNewUser(context.Background(), username, password, telefone, name, email)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, userId)
}

func EditUserPermission(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	permission := json_map["permission"].(string) //está dando erro quando tenta pegar o "email" e ele não existe.
	userId := json_map["userId"].(string)         //está dando erro quando tenta pegar o "email" e ele não existe.

	if userId == "" {
		c.String(400, "Edit User Permission Error: userId not find")
		return
	}

	if permission == "" {
		c.String(400, "Edit User Permission Error: permission not find")
		return
	}

	//ver se permission é um numero valido (0,1,2,3)

	err = service.GetInstanceUser().EditPermissionUser(context.Background(), userId, permission)
	if err != nil {
		c.String(403, err.Error())
		return
	}

	c.String(http.StatusOK, "Ok")
}

func GetInformationByUserId(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	userId := json_map["userId"].(string) //está dando erro quando tenta pegar o "email" e ele não existe.

	result, err := service.GetInstanceUser().GetInformation(context.Background(), userId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}

func GetCarsUserCanCreate(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	userId := json_map["userId"].(string)

	if userId == "" {
		c.String(400, "Get cars can create: userId not exist")
		return
	}

	result, err := service.GetInstanceUser().GetCarsUserCanCreate(context.Background(), userId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}
