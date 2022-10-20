//This example uses the ORM jet
package client

import (
	"database/sql"

	"github.com/Eli15x/search-car/src/models"
	"github.com/go-sql-driver/mysql"

	"fmt"
	"sync"
)

var (
	instanceClient CommandClient
	onceClient     sync.Once
)

type CommandClient interface {
	Connect() error
	DeleteParameter(commandSql string, interfaceSql interface{}) error
	SelectOneParameter(commandSql string, interfaceSql interface{}) (string, error)
	SelectTwoParameter(commandSql string, interfaceSql interface{}, interfaceSql_second interface{}) (string, error)
	InsertUser(sqlStatement string, user *models.User) error
	InsertCar(sqlStatement string, car *models.Car) error
	UpdatePermission(sqlStatement string, user *models.User) error
	SelectUserInformation(commandSql string, interfaceSql interface{}) (map[string]interface{}, error)
	SelectUserAllInformation(commandSql string, interfaceSql interface{}) (*models.User, error)
	SelectCarAllInformation(commandSql string, interfaceSql interface{}) ([]*models.Car, error)
	SelectCarInformation(commandSql string, interfaceSql interface{}) ([]map[string]interface{}, error)
	UpdateCars(sqlStatement string, user *models.User) error
	UpdateValidation(sqlStatement string, validation *models.Validacao) error
	InsertValidation(sqlStatement string, validation *models.Validacao) error
	SelectCountCar(commandSql string, interfaceSql interface{}) (int, error)
}

type clientSql struct {
	db *sql.DB
}

func GetInstance() CommandClient {
	onceClient.Do(func() {
		instanceClient = &clientSql{}
	})
	return instanceClient
}

func (c *clientSql) Connect() error {
	//Make sure you setup the ELEPHANTSQL_URL to be a uri, e.g. 'postgres://user:pass@host/db?options'
	cfg := mysql.Config{
		User:                 "thiago",
		Passwd:               "P@1089ppp",
		Net:                  "tcp",
		Addr:                 "busca-tesouro.mysql.uhserver.com",
		DBName:               "busca_tesouro",
		AllowNativePasswords: true,
	}

	var err error
	c.db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err.Error())
	}

	return nil
}

func (c *clientSql) DeleteParameter(commandSql string, interfaceSql interface{}) error {
	_, err := c.db.Exec(commandSql, interfaceSql)
	if err != nil {
		return err
	}

	return nil
}
func (c *clientSql) SelectOneParameter(commandSql string, interfaceSql interface{}) (string, error) {
	var result string
	row := c.db.QueryRow(commandSql, interfaceSql)
	err := row.Scan(&result)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			return result, nil
		}
	}
	return result, err
}

func (c *clientSql) SelectTwoParameter(commandSql string, interfaceSql interface{}, interfaceSql_second interface{}) (string, error) {
	var result string
	row := c.db.QueryRow(commandSql, interfaceSql, interfaceSql_second)
	err := row.Scan(&result)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			return "", nil
		} else {
			return "", err
		}
	}
	return result, nil
}

func (c *clientSql) SelectUserInformation(commandSql string, interfaceSql interface{}) (map[string]interface{}, error) {
	var userId string
	var name string
	var permission string

	var userModel map[string]interface{}

	row := c.db.QueryRow(commandSql, interfaceSql)
	err := row.Scan(&name, &permission, &userId)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			return userModel, nil
		} else {
			return userModel, err
		}
	}

	userModel = map[string]interface{}{
		"UserId":     userId,
		"Name":       name,
		"Permission": permission,
	}

	//ver se o retorno do data é realmente oque desejo no caso do select.
	return userModel, nil
}

func (c *clientSql) SelectUserAllInformation(commandSql string, interfaceSql interface{}) (*models.User, error) {
	var userId string
	var username string
	var telefone string
	var name string
	var email string
	var password string
	var permission string
	var cars int

	var userModel *models.User

	row := c.db.QueryRow(commandSql, interfaceSql)
	err := row.Scan(&userId, &username, &password, &telefone, &name, &email, &permission, &cars)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			return userModel, nil
		} else {
			return userModel, err
		}
	}

	userModel = &models.User{
		UserId:     userId,
		Name:       name,
		Username:   username,
		Email:      email,
		Password:   password,
		Telefone:   telefone,
		Permission: permission,
		Cars:       cars,
	}

	//ver se o retorno do data é realmente oque desejo no caso do select.
	return userModel, nil
}

func (c *clientSql) SelectCarAllInformation(commandSql string, interfaceSql interface{}) ([]*models.Car, error) {

	var carModel []*models.Car

	rows, err := c.db.Query(commandSql, interfaceSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var Placa string
		var Cor string
		var Nome string
		var UserId string
		var Municipio string
		var State string
		var MarcaEModelo string
		var Chassi string
		var Renavam string
		var AnoDoCarro string

		err = rows.Scan(&Placa, &UserId, &Renavam, &State, &MarcaEModelo, &Municipio, &AnoDoCarro, &Cor, &Chassi, &Nome)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Zero rows found")
				return nil, nil
			} else {
				return nil, err
			}
		}

		var newCar = &models.Car{
			UserId:       UserId,
			Nome:         Nome,
			Cor:          Cor,
			Municipio:    Municipio,
			State:        State,
			MarcaEModelo: MarcaEModelo,
			AnoDoCarro:   AnoDoCarro,
			Placa:        Placa,
		}

		carModel = append(carModel, newCar)
	}

	//ver se o retorno do data é realmente oque desejo no caso do select.
	return carModel, nil
}

func (c *clientSql) SelectCarInformation(commandSql string, interfaceSql interface{}) ([]map[string]interface{}, error) {

	var carModel []map[string]interface{}

	rows, err := c.db.Query(commandSql, interfaceSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var Placa string
		var Cor string
		var Nome string
		var UserId string
		var Municipio string
		var State string
		var MarcaEModelo string
		var Chassi string
		var Renavam string
		var AnoDoCarro string

		err = rows.Scan(&Placa, &UserId, &Renavam, &State, &MarcaEModelo, &Municipio, &AnoDoCarro, &Cor, &Chassi, &Nome)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Zero rows found")
				return nil, nil
			} else {
				return nil, err
			}
		}

		var newCar = map[string]interface{}{
			"UserId":       UserId,
			"Nome":         Nome,
			"Cor":          Cor,
			"Municipio":    Municipio,
			"State":        State,
			"MarcaEModelo": MarcaEModelo,
			"AnoDoCarro":   AnoDoCarro,
			"Placa":        Placa,
		}

		carModel = append(carModel, newCar)
	}

	//ver se o retorno do data é realmente oque desejo no caso do select.
	return carModel, nil
}

func (c *clientSql) SelectCountCar(commandSql string, interfaceSql interface{}) (int, error) {

	var countCar = 0

	rows, err := c.db.Query(commandSql, interfaceSql)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		countCar++
	}

	return countCar, nil
}

func (c *clientSql) InsertUser(sqlStatement string, user *models.User) error {

	_, err := c.db.Exec(sqlStatement, user.UserId, user.Username, user.Password, user.Telefone, user.Name, user.Email, user.Permission, user.Cars)

	if err != nil {
		return err
	}

	return nil
}

func (c *clientSql) InsertCar(sqlStatement string, car *models.Car) error {
	_, err := c.db.Exec(sqlStatement, car.Placa, car.UserId, car.State, car.MarcaEModelo, car.Municipio, car.Nome, car.Cor, car.Renavam, car.Chassi, car.AnoDoCarro)

	if err != nil {
		return err
	}

	return nil
}

func (c *clientSql) UpdatePermission(sqlStatement string, user *models.User) error {

	_, err := c.db.Exec(sqlStatement, user.Permission, user.UserId)

	if err != nil {
		return err
	}

	return nil
}

func (c *clientSql) UpdateCars(sqlStatement string, user *models.User) error {

	_, err := c.db.Exec(sqlStatement, user.Cars, user.UserId)

	if err != nil {
		return err
	}

	return nil
}

func (c *clientSql) UpdateValidation(sqlStatement string, validation *models.Validacao) error {

	_, err := c.db.Exec(sqlStatement, validation.Number, validation.UserId)

	if err != nil {
		return err
	}

	return nil
}

func (c *clientSql) InsertValidation(sqlStatement string, validation *models.Validacao) error {

	_, err := c.db.Exec(sqlStatement, validation.UserId, validation.Number)

	if err != nil {
		return err
	}

	return nil
}

//criar funcao de deletar linha de tabela.
