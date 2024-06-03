package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Firstname   string              `json:"firstname" bson:"firstname,omitempty"`
	Lastname    string              `json:"lastname" bson:"lastname,omitempty"`
	Email       string              `json:"email" bson:"email,omitempty"`
	Status      string              `json:"status" bson:"status,omitempty"`
	Tel         string              `json:"tel" bson:"tel,omitempty"`
	BirthDate   *time.Time          `json:"birthDate" bson:"birthDate,omitempty"`
	CreatedDate time.Time           `json:"createdDate" bson:"createdDate"`
	UpdatedDate time.Time           `json:"updatedDate" bson:"updatedDate"`
}

var USER_STATUS = []string{"registered", "closing", "closed"}

type UserResponse struct {
	Id          string    `json:"id"`
	Status      bool      `json:"status"`
	UpdatedDate time.Time `json:"updatedDate"`
}
