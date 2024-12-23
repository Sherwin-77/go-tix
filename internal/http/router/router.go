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
	eventApprovalHandler handler.EventApprovalHandler,
	saleInvoiceHandler handler.SaleInvoiceHandler,
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
			Method:  http.MethodGet,
			Path:    "/profile",
			Handler: userHandler.ShowProfile,
			Middlewares: []echo.MiddlewareFunc{
				authMiddleware.Authenticated,
			},
		},
		{
			Method:  http.MethodPatch,
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
		{
			Method:  http.MethodPost,
			Path:    "/register-event",
			Handler: eventHandler.RegisterEvent,
			Middlewares: []echo.MiddlewareFunc{
				authMiddleware.Authenticated,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/events-approvals",
			Handler: eventApprovalHandler.GetUserEventApprovals,
			Middlewares: []echo.MiddlewareFunc{
				authMiddleware.Authenticated,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/events-approvals/:id",
			Handler: eventApprovalHandler.GetUserEventApprovalByID,
			Middlewares: []echo.MiddlewareFunc{
				authMiddleware.Authenticated,
				validateID,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/sale-invoices",
			Handler: saleInvoiceHandler.GetUserSaleInvoices,
			Middlewares: []echo.MiddlewareFunc{
				authMiddleware.Authenticated,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/sale-invoices/:id",
			Handler: saleInvoiceHandler.GetUserSaleInvoiceByID,
			Middlewares: []echo.MiddlewareFunc{
				authMiddleware.Authenticated,
				validateID,
			},
		},
		{
			Method:  http.MethodPost,
			Path:    "/bill",
			Handler: saleInvoiceHandler.Bill,
			Middlewares: []echo.MiddlewareFunc{
				authMiddleware.Authenticated,
			},
		},
		{
			Method:  http.MethodPost,
			Path:    "/checkout",
			Handler: saleInvoiceHandler.Checkout,
			Middlewares: []echo.MiddlewareFunc{
				authMiddleware.Authenticated,
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
	eventApprovalHandler handler.EventApprovalHandler,
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
		{
			Method:      http.MethodPost,
			Path:        "/events",
			Handler:     eventHandler.CreateEvent,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/events/:id",
			Handler: eventHandler.UpdateEvent,
			Middlewares: []echo.MiddlewareFunc{
				validateID,
			},
		},
		{
			Method:      http.MethodGet,
			Path:        "/event-approvals",
			Handler:     eventApprovalHandler.GetEventApprovals,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Method:  http.MethodGet,
			Path:    "/event-approvals/:id",
			Handler: eventApprovalHandler.GetEventApprovalByID,
			Middlewares: []echo.MiddlewareFunc{
				validateID,
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/event-approvals/:id",
			Handler: eventApprovalHandler.HandleEventApproval,
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
