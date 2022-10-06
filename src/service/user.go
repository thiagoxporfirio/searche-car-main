package service

import (
	"context"
	"errors"
	"sync"

	"github.com/Eli15x/search-car/src/client"
	"github.com/Eli15x/search-car/src/models"
	"github.com/Eli15x/search-car/src/utils"
)

var (
	instanceUser CommandUser
	onceUser     sync.Once
)

type CommandUser interface {
	CreateNewUser(ctx context.Context, username string, password string, telefone string, name string, email string) (string, error)
	ValidateUser(ctx context.Context, username string, password string) (map[string]interface{}, error)
	EditPermissionUser(ctx context.Context, userId string, permission string) error
	GetInformation(ctx context.Context, userId string) (*models.User, error)
}

type user struct{}

func GetInstanceUser() CommandUser {
	onceUser.Do(func() {
		instanceUser = &user{}
	})
	return instanceUser
}

func UserExistsEmail(email string) error {

	data, err := client.GetInstance().SelectOneParameter(`SELECT userId FROM userInfo WHERE email = ?`, email)

	if err != nil {
		return errors.New("User Exists: error try get user info in Sql")
	}

	if data != "" {
		return errors.New("User Exists: email already exists in Sql")
	}

	return nil
}

func UserExistsUsername(username string) error {

	data, err := client.GetInstance().SelectOneParameter(`SELECT * FROM userInfo WHERE username = ?`,
		username)

	if err != nil {
		return errors.New("User Exists: error try get user info in Sql")
	}

	if data != "" {
		return errors.New("User Exists: username already exists in Sql")
	}

	return nil
}

func (u *user) ValidateUser(ctx context.Context, email string, password string) (map[string]interface{}, error) {

	result, err := client.GetInstance().SelectTwoParameter(`SELECT userId FROM userInfo WHERE email = ? and password = ?`,
		email,
		password,
	)

	if result == "" {
		return nil, errors.New("Validate User: user not exists")
	}

	if err != nil {
		return nil, errors.New("Validate User: error validate info in sql")
	}

	resultUser, err := client.GetInstance().SelectUserInformation(`SELECT name,permission,userId  FROM userInfo WHERE userId = ?`, result)

	newResult, err := client.GetInstance().SelectOneParameter(`SELECT number FROM validation WHERE userId = ?`,
		result)

	if newResult != "" {
		return resultUser, errors.New("Validate User: error user in validation")
	}

	return resultUser, nil
}

func (u *user) CreateNewUser(ctx context.Context, username string, password string, telefone string, name string, email string) (string, error) {

	err := UserExistsEmail(email)
	if err != nil {
		return "", err
	}

	userId := utils.CreateCodeId()
	userModel := &models.User{
		UserId:     userId,
		Name:       name,
		Username:   username,
		Email:      email,
		Password:   password,
		Telefone:   telefone,
		Permission: "0",
		Cars:       0,
	}

	err = GetInstanceValidation().SendMail(email, userId)
	if err != nil {
		return "", err
	}

	err = client.GetInstance().InsertUser(`INSERT INTO busca_tesouro.userInfo (userId,username,password,telefone,name,email,permission,cars) VALUES (?,?,?,?,?,?,?,?)`, userModel)

	if err != nil {
		return "", errors.New("Create New User: problem to insert into Sql")
	}

	return userId, nil
}

func (u *user) EditPermissionUser(ctx context.Context, userId string, permission string) error {

	userModel := &models.User{
		UserId:     userId,
		Permission: permission,
	}

	err := client.GetInstance().UpdatePermission(`UPDATE userInfo SET permission = ? WHERE userId = ?`, userModel)
	if err != nil {
		return errors.New("Update Permission User: problem to update into Sql")
	}

	return nil
}

func (u *user) EditCarsUser(ctx context.Context, userId string, cars int) error {

	userModel := &models.User{
		UserId: userId,
		Cars:   cars,
	}

	err := client.GetInstance().UpdateCars(`UPDATE userInfo SET cars = ? WHERE userId = ?`, userModel)
	if err != nil {
		return errors.New("Update Permission User: problem to update into Sql")
	}

	return nil
}

func (u *user) GetInformation(ctx context.Context, userId string) (*models.User, error) {

	result, err := client.GetInstance().SelectUserAllInformation(`SELECT * FROM userInfo WHERE userId = ?`, userId)

	if result == nil {
		return nil, errors.New("Get Information: user not exists")
	}

	if err != nil {
		return nil, errors.New("Get Information: error validate info in sql")
	}

	return result, nil
}
