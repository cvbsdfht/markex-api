package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/markex-api/pkg/core"
	"github.com/markex-api/pkg/modules/users/service"
	"github.com/markex-api/pkg/tools/handler"
)

type userRouteHandler struct {
	app         *fiber.App
	core        *core.CoreRegistry
	userService service.IUserService
}

func NewUserRouteHandler(app *fiber.App, core *core.CoreRegistry, userService service.IUserService) *userRouteHandler {
	return &userRouteHandler{app, core, userService}
}

func (h *userRouteHandler) Init() {
	router := h.app

	api := router.Group("/api/user", h.middleware)
	api.Get("/", h.getUserList)
	api.Get("/:id", h.getUserById)
}

func (h *userRouteHandler) middleware(c *fiber.Ctx) error {
	return c.Next()
}

func (h *userRouteHandler) getUserList(c *fiber.Ctx) error {
	return handler.Handler(c, func() (interface{}, error) {
		return h.userService.GetUserList()
	})
}

func (h *userRouteHandler) getUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	return handler.Handler(c, func() (interface{}, error) {
		return h.userService.GetUserById(id)
	})
}
