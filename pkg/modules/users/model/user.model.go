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

var USER_STATUS = []string{"registered", "closing", "closed"}

type UserResponse struct {
	Id          string    `json:"id"`
	Status      bool      `json:"status"`
	UpdatedDate time.Time `json:"updatedDate"`
}
