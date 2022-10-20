package main

import (
	//"time"
	//"context"

	"github.com/Eli15x/search-car/src/client"
	"github.com/Eli15x/search-car/src/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func main() {
	//Context
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancel()
	if err := client.GetInstance().Connect(); err != nil {
		log.Infof("[ElephantSql] problem to Connect : %s \n", err, "")
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.POST("/login", handlers.ValidateUser)
	router.POST("/cadastro", handlers.CreateUser)
	router.POST("/user/change-permission", handlers.EditUserPermission)
	router.POST("/user/getInformation", handlers.GetInformationByUserId)
	router.POST("/car/create", handlers.CreateCar)
	router.POST("/car/placa", handlers.GetUserBycar)
	router.POST("/car/state", handlers.GetCarsByState)
	router.POST("/car/carsCanCreate", handlers.GetCarsUserCanCreate)
	router.POST("/car/cars", handlers.GetCarsByUserId)
	router.POST("/validation", handlers.ValidateNumber)
	router.POST("/sendInformation", handlers.SendEmailCar)

	router.Run()
}
