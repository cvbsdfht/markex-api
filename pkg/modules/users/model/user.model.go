package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Firstname   string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname    string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email       string             `json:"email" bson:"email,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	Tel         string             `json:"tel,omitempty" bson:"tel,omitempty"`
	BirthDate   *time.Time         `json:"birthDate,omitempty" bson:"birthDate,omitempty"`
	CreatedDate time.Time          `json:"createdDate" bson:"createdDate,omitempty"`
	UpdatedDate time.Time          `json:"updatedDate" bson:"updatedDate,omitempty"`
	DeletedDate *time.Time         `json:"deletedDate,omitempty" bson:"deletedDate,omitempty"`
}

type UserRequest struct {
	Id        string `json:"id" validate:"mongodb"`
	Firstname string `json:"firstname" validate:"name"`
	Lastname  string `json:"lastname" validate:"name"`
	Email     string `json:"email" validate:"required,email"`
	Tel       string `json:"tel"`
	BirthDate string `json:"birthDate"`
}

type UserResponse struct {
	Status bool  `json:"status"`
	Data   *User `json:"data"`
}

type UserListResponse struct {
	Status bool    `json:"status"`
	Data   *[]User `json:"data"`
}

type UserLoginRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type UserLoginResponse struct {
	Email string
	Token string
}

const (
	USER_STATUS_REGISTERED = "registered"
	USER_STATUS_CLOSING    = "closing"
	USER_STATUS_CLOSED     = "closed"
)
