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

func SendEmailCar(c *gin.Context) {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	nome := json_map["nome"].(string)
	email := json_map["email"].(string)
	placa := json_map["placa"].(string)
	marca := json_map["marca"].(string)
	municipio := json_map["municipio"].(string)
	estado := json_map["estado"].(string)
	localizacao := json_map["localizacao"].(string)
	foto1 := json_map["foto1"].(string)
	foto2 := json_map["foto2"].(string)
	foto3 := json_map["foto3"].(string)
	foto4 := json_map["foto4"].(string)


	if nome == "" {
		c.String(http.StatusBadRequest, "Send Email Car Error: nome not find")
		return
	}

	if email == "" {
		c.String(http.StatusBadRequest, "Send Email Car Error: email not find")
		return
	}

	if placa == "" {
		c.String(http.StatusBadRequest, "Send Email Car Error: placa not find")
		return
	}

	if marca == "" {
		c.String(http.StatusBadRequest, "Send Email Car Error: marca not find")
		return
	}

	if municipio == "" {
		c.String(http.StatusBadRequest, "Send Email Car Error: municipio not find")
		return
	}

	if estado == "" {
		c.String(http.StatusBadRequest, "Send Email Car Error: estado not find")
		return
	}

	if localizacao == "" {
		c.String(http.StatusBadRequest, "Send Email Car Error: localizacao not find")
		return
	}

	err = service.GetInstanceValidation().SendEmailInformation(nome, email, localizacao, placa, marca, municipio, estado, foto1, foto2, foto3, foto4)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}
