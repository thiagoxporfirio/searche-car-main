package service

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"

	gomail "gopkg.in/mail.v2"

	"github.com/Eli15x/search-car/src/client"
	"github.com/Eli15x/search-car/src/models"
)

var (
	instanceValidation CommandValidation
	onceValidation     sync.Once
)

type CommandValidation interface {
	ValidateNumber(userId string, number string) error
	SendMail(toEmail string, userId string) error
}

type validation struct{}

func GetInstanceValidation() CommandValidation {
	onceValidation.Do(func() {
		instanceValidation = &validation{}
	})
	return instanceValidation
}

func (v *validation) SendMail(toEmail string, userId string) error {

	number := strconv.Itoa(rand.Intn(9999))

	emailSender := gomail.NewMessage()
	body := "<html><body><h1>Olá seja bem vindo ao busca tesouro direto!</h1><br/> Seu código de autenticação é: " + number + "<br/> <a href='https://proj-busca-tesouro.vercel.app/cadastroUser/validateEmail.html' alt='tela'>clique aqui</a></body></html>"

	emailSender.SetHeader("From", "contato@buscatesouro.com.br")
	emailSender.SetHeader("To", toEmail)
	emailSender.SetHeader("Subject", "Seu cógio de acesso chegou!")
	emailSender.SetBody("text/html", body)

	userValidator := &models.Validacao{
		UserId: userId,
		Number: number,
	}

	err := userHaveValidation(userId)
	if err == errors.New("Get validation: user have validation") {
		err = client.GetInstance().UpdateValidation(`UPDATE validation SET number = ? WHERE userId = ?`, userValidator)
		if err != nil {
			return errors.New("Problem validation: error sql")
		}
	} else {
		err = client.GetInstance().InsertValidation(`INSERT INTO busca_tesouro.validation (userId,number) VALUES (?,?)`, userValidator)
		if err != nil {
			return errors.New("Problem validation: error sql")
		}
	}

	send := gomail.NewDialer("smtps.uhserver.com", 465, "contato@buscatesouro.com.br", "Busc@tesouro2022")

	if err := send.DialAndSend(emailSender); err != nil {
		return err
	}

	return nil
}

func userHaveValidation(userId string) error {
	result, err := client.GetInstance().SelectUserAllInformation(`SELECT * FROM validation WHERE userId = ?`, userId)

	if result == nil {
		return nil
	}

	if err != nil {
		return errors.New("Get Validation: error validate info in sql")
	}

	return errors.New("Get validation: user have validation")

}

func (v *validation) ValidateNumber(userId string, number string) error {
	result, err := client.GetInstance().SelectTwoParameter(`SELECT number FROM validation WHERE userId = ? and number = ?`,
		userId,
		number,
	)

	if result == "" {
		return errors.New("Validate User: number not correct")
	}

	if err != nil {
		return errors.New("Validate User: error validate info in sql")
	}

	err = client.GetInstance().DeleteParameter(`DELETE FROM validation WHERE userId = ?`, userId)
	if err != nil {
		return errors.New("Validate User: error get data")
	}

	return nil
}
