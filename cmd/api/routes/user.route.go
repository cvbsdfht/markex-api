package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/markex-api/pkg/core"
	userModel "github.com/markex-api/pkg/modules/users/model"
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

	api := router.Group("/api/user", h.validation)
	api.Get("/", h.getUserList)
	api.Get("/:id", h.getUserById)

	api.Post("/register", h.registerUser)
	api.Post("/update", h.updateUser)
	api.Delete("/:id", h.closingUser)
}

func (h *userRouteHandler) validation(c *fiber.Ctx) error {
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

func (h *userRouteHandler) registerUser(c *fiber.Ctx) error {
	request := &userModel.UserRequest{}
	return handler.HandlerWithBody(c, request, func() (interface{}, error) {
		return h.userService.Create(request)
	})
}

func (h *userRouteHandler) updateUser(c *fiber.Ctx) error {
	request := &userModel.UserRequest{}
	return handler.HandlerWithBody(c, request, func() (interface{}, error) {
		return h.userService.Update(request)
	})
}

func (h *userRouteHandler) closingUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return handler.Handler(c, func() (interface{}, error) {
		return h.userService.Closing(id)
	})
}