package routes

import "github.com/gofiber/fiber/v2"

type healthRouteHandler struct {
	app *fiber.App
}

func NewHealthRouteHandler(app *fiber.App) *healthRouteHandler {
	return &healthRouteHandler{app}
}

func (h *healthRouteHandler) Init() {
	h.app.Get("/health", h.health)
}

func (h *healthRouteHandler) health(c *fiber.Ctx) error {
	return c.SendString("Healthy")
}
