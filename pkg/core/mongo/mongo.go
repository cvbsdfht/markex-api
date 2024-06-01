package mongo

import (
	"context"

	"github.com/markex-api/pkg/core/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnect(uri string, log logger.Logger) *mongo.Client {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Error(err)
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		log.Error(err)
		panic(err)
	}
	log.Info("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}
