package main

import (
	"fmt"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/markex-api/cmd/api/routes"
	"github.com/markex-api/pkg/core"
	"github.com/markex-api/pkg/core/config"
	"github.com/markex-api/pkg/core/logger"
	"github.com/markex-api/pkg/core/mongo"
	"github.com/markex-api/pkg/core/redis"
	"github.com/markex-api/pkg/modules"
	"github.com/markex-api/pkg/modules/users/repository"
	"github.com/markex-api/pkg/modules/users/service"
)

func main() {
	// Load config
	config := config.NewConfig("config/api/config.local.yml")
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

	// Create fiber
	app := fiber.New(fiber.Config{
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy")
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString(fmt.Sprintf("Hello %v", name))
	})

	// Core
	coreRegistry := &core.CoreRegistry{
		Configuration: cfg,
		Logger:        log,
		Mongo:         mongoClient,
		Redis:         rdb,
	}

	// Repository
	userRepository := repository.NewUserRepository(mongoDb)

	repositoryRegistry := &modules.RepositoryRegistry{
		UserRepository: userRepository,
	}

	// Service
	userService := service.NewUserService(coreRegistry, repositoryRegistry)
	_ = userService

	// Route
	routes.NewUserRouteHandler(app, coreRegistry, userService).Init()

	// Listen
	app.Listen(fmt.Sprintf(":%v", cfg.App.Port))
}
