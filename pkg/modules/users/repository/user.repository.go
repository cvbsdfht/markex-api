package repository

import (
	"context"
	"time"

	userModel "github.com/markex-api/pkg/modules/users/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Port
type IUserRepository interface {
	GetAll() (*[]userModel.User, error)
	GetById(id primitive.ObjectID) (*userModel.User, error)
	Create(user *userModel.User) (interface{}, error)
}

// Adaptor
type userRepository struct {
	UserCollection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) IUserRepository {
	userCollection := database.Collection("users")
	return &userRepository{UserCollection: userCollection}
}

func (r *userRepository) GetAll() (*[]userModel.User, error) {
	users := &[]userModel.User{}

	filter := bson.M{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filterCursor, err := r.UserCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err = filterCursor.All(ctx, users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetById(id primitive.ObjectID) (*userModel.User, error) {
	user := &userModel.User{}

	filter := bson.M{"_id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.UserCollection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Create(user *userModel.User) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}
