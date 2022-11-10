package handlers

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/Eli15x/search-car/src/service"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func CreateCar(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	municipio := json_map["municipio"].(string)
	marcaEModelo := json_map["marcaEModelo"].(string)
	nome := json_map["nome"].(string)
	cor := json_map["cor"].(string)
	placa := json_map["placa"].(string)
	state := json_map["state"].(string) //está dando erro quando tenta pegar o "email" e ele não existe.
	userId := json_map["userId"].(string)
	renavam := json_map["renavam"].(string)
	chassi := json_map["chassi"].(string)
	anoDoCarro := json_map["anoDoCarro"].(string)

	if marcaEModelo == "" {
		c.String(http.StatusBadRequest, "Create User Error: marcaEModelo do carro not find")
		return
	}

	if municipio == "" {
		c.String(http.StatusBadRequest, "Create User Error: municipio not find")
		return
	}

	if nome == "" {
		c.String(http.StatusBadRequest, "Create User Error: nome not find")
		return
	}

	if cor == "" {
		c.String(http.StatusBadRequest, "Create User Error: cor not find")
		return
	}

	if state == "" {
		c.String(http.StatusBadRequest, "Create User Error: state not find")
		return
	}

	if userId == "" {
		c.String(400, "Create User Error: userId not find")
		return
	}

	if anoDoCarro == "" {
		c.String(400, "Create User Error: anoDoCarro not find")
		return
	}

	if chassi == "" {
		c.String(400, "Create User Error: chassi not find")
		return
	}

	if renavam == "" {
		c.String(400, "Create User Error: renavam not find")
		return
	}

	err = service.GetInstanceCar().CreateNewCar(context.Background(), placa, state, userId, marcaEModelo, municipio, nome, cor, anoDoCarro, renavam, chassi)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "ok")
}

func GetUserBycar(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	placa := json_map["placa"].(string)

	if placa == "" {
		c.String(400, "Get User By car Error: placa not find")
		return
	}

	result, err := service.GetInstanceCar().GetInformationUser(context.Background(), placa)
	if err != nil {
		c.String(403, err.Error())
		return
	}

	c.String(http.StatusOK, result)
}

func GetCar(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	placa := json_map["placa"].(string)

	if placa == "" {
		c.String(400, "Get Car Error: placa not find")
		return
	}

	result, err := service.GetInstanceCar().GetCar(context.Background(), placa)
	if err != nil {
		c.String(403, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetCarsByState(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	state := json_map["state"].(string)

	if state == "" {
		c.String(400, "Get cars by State: state not exist")
		return
	}

	result, err := service.GetInstanceCar().GetInformationCarByState(context.Background(), state)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}

func GetCarsByUserId(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	userId := json_map["userId"].(string)

	if userId == "" {
		c.String(400, "Get cars by UserId: userId not exist")
		return
	}

	result, err := service.GetInstanceCar().GetCarsByUserId(context.Background(), userId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}
