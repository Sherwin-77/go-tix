package router

import (
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-tix/internal/http/handler"
	"github.com/sherwin-77/go-tix/pkg/route"
	"net/http"
)

func WebhookRoutes(webhookHandler handler.WebhookHandler) ([]route.Route, []echo.MiddlewareFunc) {
	return []route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/midtrans",
			Handler: webhookHandler.HandleMidtransNotification,
		},
	}, nil
}
