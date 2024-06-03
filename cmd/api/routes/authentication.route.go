package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/markex-api/pkg/core"
	userModel "github.com/markex-api/pkg/modules/users/model"
	"github.com/markex-api/pkg/modules/users/service"
	"github.com/markex-api/pkg/tools/handler"
)

type authenticationRouteHandler struct {
	app         *fiber.App
	core        *core.CoreRegistry
	userService service.IUserService
}

func NewAuthenticationRouteHandler(app *fiber.App, core *core.CoreRegistry, userService service.IUserService) *authenticationRouteHandler {
	return &authenticationRouteHandler{app, core, userService}
}

func (h *authenticationRouteHandler) Init() {
	router := h.app

	api := router.Group("/api/login", h.validation)
	api.Post("/", h.login)
}

func (h *authenticationRouteHandler) validation(c *fiber.Ctx) error {
	return c.Next()
}

func (h *authenticationRouteHandler) login(c *fiber.Ctx) error {
	request := &userModel.UserLoginRequest{}
	return handler.HandlerWithBody(c, request, func() (interface{}, error) {
		return h.userService.Login(request)
	})
}
