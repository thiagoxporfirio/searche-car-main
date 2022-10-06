package models

type Validacao struct {
	UserId string `json:"userId,omitempty" bson:"userId,omitempty"`
	Number string `json:"number,omitempty" bson:"-"`
}
