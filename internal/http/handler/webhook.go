package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-tix/internal/service"
	"github.com/sherwin-77/go-tix/pkg/response"
	"net/http"
)

type WebhookHandler interface {
	HandleMidtransNotification(ctx echo.Context) error
}

type webhookHandler struct {
	webhookService service.WebhookService
}

func NewWebhookHandler(webhookService service.WebhookService) WebhookHandler {
	return &webhookHandler{
		webhookService: webhookService,
	}
}

func (h *webhookHandler) HandleMidtransNotification(ctx echo.Context) error {
	payload := make(map[string]interface{})

	if err := json.NewDecoder(ctx.Request().Body).Decode(&payload); err != nil {
		ctx.Logger().Error(err)
		return err
	}

	if err := h.webhookService.HandleMidtransNotification(ctx.Request().Context(), payload); err != nil {
		ctx.Logger().Error(err)
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", nil, nil))
}
