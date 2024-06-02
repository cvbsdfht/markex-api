package core

import (
	"github.com/go-redis/redis"
	"github.com/markex-api/pkg/core/config"
	"github.com/markex-api/pkg/core/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type CoreRegistry struct {
	Configuration *config.Configuration
	Logger        logger.Logger
	Mongo         *mongo.Client
	Redis         *redis.Client
}
