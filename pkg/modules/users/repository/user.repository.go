package repository

import (
	"context"
	"time"

	userModel "github.com/markex-api/pkg/modules/users/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Port
type IUserRepository interface {
	GetAll() (*[]userModel.User, error)
	GetById(id primitive.ObjectID) (*userModel.User, error)
	Upsert(user *userModel.User) (*userModel.User, error)
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

	filter := bson.M{
		"status": "registered",
	}
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

	filter := bson.M{
		"_id": id,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.UserCollection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Upsert(user *userModel.User) (*userModel.User, error) {
	filter := bson.M{
		"_id":    user.Id,
		"status": "registered",
	}
	update := bson.D{
		{Key: "$set", Value: user},
	}
	option := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := &userModel.User{}
	err := r.UserCollection.FindOneAndUpdate(ctx, filter, update, option).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
