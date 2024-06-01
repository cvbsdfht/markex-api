package main

import (
	"log"

	"github.com/markex-api/pkg/core/config"
	"github.com/markex-api/pkg/core/logger"
	"github.com/markex-api/pkg/core/mongo"
	"github.com/markex-api/pkg/core/redis"
)

func main() {
	// Load config
	config := config.NewConfig("config/api/config.yml")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Load logger
	log := logger.NewLogger(&logger.Options{
		FilePath: cfg.Log.FilePath,
		Level:    cfg.Log.Level,
		Format:   cfg.Log.Format,
		ProdMode: cfg.App.ProdMode,
	})

	// Connect mongodb
	mongoClient := mongo.NewConnect(cfg.Mongo.Uri, log)
	mongoDb := mongoClient.Database(cfg.Mongo.Database)
	_ = mongoDb

	// Connect redis
	rdb := redis.NewConnect(cfg, log)
	_ = rdb
}
