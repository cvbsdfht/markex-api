package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/markex-api/pkg/core"
	"github.com/markex-api/pkg/modules/users/service"
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
	users, err := h.userService.GetUserList()
	if err != nil {
		h.core.Logger.Error(err)
		return c.SendStatus(400)
	}
	return c.JSON(users)
}

func (h *userRouteHandler) getUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userService.GetUserById(id)
	if err != nil {
		h.core.Logger.Error(err)
		return c.SendStatus(400)
	}
	return c.JSON(user)
}
