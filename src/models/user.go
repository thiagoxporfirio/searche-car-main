package models

type User struct {
	UserId     string `json:"userId,omitempty" bson:"userId,omitempty"`
	Username   string `json:"username,omitempty" bson:"-"`
	Telefone   string `json:"telefone,omitempty" bson:"-"`
	Name       string `json:"name,omitempty" bson:"-"`
	Email      string `json:"email,omitempty" bson:"-"`
	Password   string `json:"password,omitempty" bson:"-"`
	Permission string `json:"permission,omitempty" bson:"-"`
	Cars       int    `json:"cars,omitempty" bson:"cars"`
}
