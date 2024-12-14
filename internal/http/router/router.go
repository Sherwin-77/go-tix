package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-tix/internal/http/handler"
	"github.com/sherwin-77/go-tix/internal/http/middlewares"
	"github.com/sherwin-77/go-tix/pkg/route"
)

func UserRoutes(
	userHandler handler.UserHandler,
	eventHandler handler.EventHandler,
	authMiddleware middlewares.AuthMiddleware,
	middleware middlewares.Middleware,
) ([]route.Route, []echo.MiddlewareFunc) {
	validateID := middleware.ValidateUUID([]string{"id"})

	routes := []route.Route{
		{
			Method:      http.MethodPost,
			Path:        "/register",
			Handler:     userHandler.Register,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Method:      http.MethodPost,
			Path:        "/login",
			Handler:     userHandler.Login,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Method:  http.MethodPut,
			Path:    "/profile",
			Handler: userHandler.EditProfile,
			Middlewares: []echo.MiddlewareFunc{
				authMiddleware.Authenticated,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/events",
			Handler: eventHandler.GetUserEvents,
			Middlewares: []echo.MiddlewareFunc{
				authMiddleware.Authenticated,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/events/:id",
			Handler: eventHandler.GetUserEventByID,
			Middlewares: []echo.MiddlewareFunc{
				authMiddleware.Authenticated,
				validateID,
			},
		},
	}

	var middlewareFuncs []echo.MiddlewareFunc

	return routes, middlewareFuncs
}

func AdminRoutes(
	userHandler handler.UserHandler,
	roleHandler handler.RoleHandler,
	eventHandler handler.EventHandler,
	middleware middlewares.Middleware,
	authMiddleware middlewares.AuthMiddleware,
) ([]route.Route, []echo.MiddlewareFunc) {
	validateID := middleware.ValidateUUID([]string{"id"})

	routes := []route.Route{
		{
			Method:      http.MethodGet,
			Path:        "/users",
			Handler:     userHandler.GetUsers,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Method:      http.MethodPost,
			Path:        "/users",
			Handler:     userHandler.CreateUser,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/:id",
			Handler: userHandler.UpdateUser,
			Middlewares: []echo.MiddlewareFunc{
				validateID,
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/:id/role",
			Handler: userHandler.ChangeRole,
			Middlewares: []echo.MiddlewareFunc{
				validateID,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/:id",
			Handler: userHandler.GetUserByID,
			Middlewares: []echo.MiddlewareFunc{
				validateID,
			},
		},
		{
			Method:      http.MethodGet,
			Path:        "/roles",
			Handler:     roleHandler.GetRoles,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Method:      http.MethodPost,
			Path:        "/roles",
			Handler:     roleHandler.CreateRole,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Method:  http.MethodGet,
			Path:    "/roles/:id",
			Handler: roleHandler.GetRoleByID,
			Middlewares: []echo.MiddlewareFunc{
				validateID,
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/roles/:id",
			Handler: roleHandler.UpdateRole,
			Middlewares: []echo.MiddlewareFunc{
				validateID,
			},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/roles/:id",
			Handler: roleHandler.DeleteRole,
			Middlewares: []echo.MiddlewareFunc{
				validateID,
			},
		},
		{
			Method:      http.MethodGet,
			Path:        "/events",
			Handler:     eventHandler.GetEvents,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Method:  http.MethodGet,
			Path:    "/events/:id",
			Handler: eventHandler.GetEventByID,
			Middlewares: []echo.MiddlewareFunc{
				validateID,
			},
		},
	}

	middlewareFuncs := []echo.MiddlewareFunc{
		authMiddleware.Authenticated,
		authMiddleware.AuthLevel(2),
	}

	return routes, middlewareFuncs

}
