package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-tix/internal/http/dto"
	"github.com/sherwin-77/go-tix/internal/service"
	"github.com/sherwin-77/go-tix/pkg/response"
	"net/http"
)

type EventApprovalHandler struct {
	eventApprovalService service.EventApprovalService
}

func NewEventApprovalHandler(eventApprovalService service.EventApprovalService) EventApprovalHandler {
	return EventApprovalHandler{eventApprovalService: eventApprovalService}
}

/* -------------------------------------------------------------------------- */
/*                                Admin Handler                               */
/* -------------------------------------------------------------------------- */

// GetEventApprovals
//
//	@Summary	Get Event Approvals
//	@Tags		[Admin] Event Approval
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.Response{data=[]dto.EventApprovalResponse}
//	@Router		/admin/event-approvals [get]
func (e *EventApprovalHandler) GetEventApprovals(ctx echo.Context) error {
	eventApprovals, meta, err := e.eventApprovalService.GetEventApprovals(ctx.Request().Context(), ctx.QueryParams())

	if err != nil {
		return err
	}

	eventApprovalResponse := dto.NewEventApprovalsResponse(eventApprovals)

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", eventApprovalResponse, meta))
}

// GetEventApprovalByID
//
//	@Summary	Get Event Approval By ID
//	@Tags		[Admin] Event Approval
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"The Event Approval ID"
//	@Success	200	{object}	response.Response{data=dto.EventApprovalResponse}
//	@Router		/admin/event-approvals/{id} [get]
func (e *EventApprovalHandler) GetEventApprovalByID(ctx echo.Context) error {
	eventApprovalID := ctx.Param("id")
	if eventApprovalID == "" {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	eventApproval, err := e.eventApprovalService.GetEventApprovalByID(ctx.Request().Context(), eventApprovalID)
	if err != nil {
		return err
	}

	eventApprovalResponse := dto.NewEventApprovalResponse(eventApproval)

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", eventApprovalResponse, nil))
}

// HandleEventApproval
//
//	@Summary	Handle Event Approval
//	@Tags		[Admin] Event Approval
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string							true	"The Event Approval ID"
//	@Param		request	body		dto.HandleEventApprovalRequest	true	"Handle Event Approval Request"
//	@Success	200		{object}	response.Response{data=dto.EventApprovalResponse}
//	@Router		/admin/event-approvals/{id} [patch]
func (e *EventApprovalHandler) HandleEventApproval(ctx echo.Context) error {
	var req dto.HandleEventApprovalRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	if err := e.eventApprovalService.HandleEventApproval(ctx.Request().Context(), req); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success! Event will automatically created if approved", nil, nil))
}

/* -------------------------------------------------------------------------- */
/*                                User Handler                                */
/* -------------------------------------------------------------------------- */
