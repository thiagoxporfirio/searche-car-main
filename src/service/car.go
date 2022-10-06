package service

import (
	"context"
	"errors"
	"sync"

	"github.com/Eli15x/search-car/src/client"
	"github.com/Eli15x/search-car/src/models"
)

var (
	instanceCar CommandCar
	onceCar     sync.Once
)

type CommandCar interface {
	CreateNewCar(ctx context.Context, placa string, state string, userId string, modeloDoCarro string, municipio string, nome string, cor string, anoDoCarro string, renavam string, chassi string) error
	GetInformationUser(ctx context.Context, placa string) (string, error)
	GetInformationCarByState(ctx context.Context, state string) ([]map[string]interface{}, error)
}

type car struct{}

func GetInstanceCar() CommandCar {
	onceCar.Do(func() {
		instanceCar = &car{}
	})
	return instanceCar
}

func CarExists(placa string) error {

	_, err := client.GetInstance().SelectOneParameter(`SELECT * FROM car WHERE placa = ?`, placa)

	if err != nil {
		return errors.New("Car Exists: placa already exists in Sql")
	}

	return nil
}

func (c *car) CreateNewCar(ctx context.Context, placa string, state string, userId string, marcaEModelo string, municipio string, nome string, cor string, anoDoCarro string, renavam string, chassi string) error {

	var user *models.User

	user, err := client.GetInstance().SelectUserAllInformation(`SELECT * FROM userInfo WHERE userId = ?`, userId)

	if user == nil {
		return errors.New("Get Information: user not exists")
	}

	if user.Permission == "0" {
		return errors.New("User don't have permission to create car")
	}

	if user.Permission == "1" && user.Cars == 1 {
		return errors.New("User don't have permission to create more cars")
	}

	if user.Permission == "2" && user.Cars == 10 {
		return errors.New("User don't have permission to create more cars")
	}

	//if 3 can create all the cars that want.
	err = CarExists(placa)
	if err != nil {
		return err
	}
	carModel := &models.Car{
		UserId:       userId,
		Placa:        placa,
		State:        state,
		MarcaEModelo: marcaEModelo,
		AnoDoCarro:   anoDoCarro,
		Cor:          cor,
		Renavam:      renavam,
		Chassi:       chassi,
		Municipio:    municipio,
		Nome:         nome,
	}
	err = client.GetInstance().InsertCar(`INSERT INTO car (placa,userId,state,marcaEModelo,municipio,nome,cor,renavam,chassi,anoDoCarro) VALUES (?,?,?,?,?,?,?,?,?,?)`, carModel)

	if err != nil {
		return errors.New("Create New Car: problem to insert into Sql")
	}

	userModel := &models.User{
		UserId: userId,
		Cars:   user.Cars + 1,
	}

	err = client.GetInstance().UpdateCars(`UPDATE userInfo SET cars = ? WHERE userId = ?`, userModel)
	if err != nil {
		return errors.New("Update Cars User: problem to update into Sql")
	}

	return nil
}

func (c *car) GetInformationUser(ctx context.Context, placa string) (string, error) {

	result, err := client.GetInstance().SelectOneParameter(`SELECT userId FROM car WHERE placa = ?`, placa)
	//fazer outra consulta.
	if result == "" {
		return "", errors.New("Get Information: user not exists")
	}

	if err != nil {
		return "", errors.New("Get Information: error validate info in sql")
	}

	return result, nil

}

func (c *car) GetInformationCarByState(ctx context.Context, state string) ([]map[string]interface{}, error) {

	result, err := client.GetInstance().SelectCarInformation(`SELECT * FROM car WHERE state = ?`, state)

	if result == nil {
		return nil, errors.New("Get Information: user not exists")
	}

	if err != nil {
		return nil, errors.New("Get Information: error validate info in sql")
	}

	return result, nil

}
