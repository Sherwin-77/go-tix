package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-tix/internal/http/dto"
	"github.com/sherwin-77/go-tix/internal/service"
	"github.com/sherwin-77/go-tix/pkg/response"
	"net/http"
)

type EventHandler struct {
	eventService service.EventService
}

func NewEventHandler(eventService service.EventService) EventHandler {
	return EventHandler{eventService: eventService}
}

/* -------------------------------------------------------------------------- */
/*                                Admin Handler                               */
/* -------------------------------------------------------------------------- */

// GetEvents
//
//	@Summary	Get All Events
//	@Tags		[Admin] Event
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.Response{data=[]dto.AdminEventListResponse}
//	@Router		/admin/events [get]
func (h *EventHandler) GetEvents(ctx echo.Context) error {
	events, meta, err := h.eventService.GetEvents(ctx.Request().Context(), ctx.QueryParams())

	if err != nil {
		return err
	}

	eventResponse := dto.NewAdminEventListResponse(events)

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", eventResponse, meta))
}

// GetEventByID
//
//	@Summary	Get Event By ID
//	@Tags		[Admin] Event
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"The Event ID"
//	@Success	200	{object}	response.Response{data=dto.AdminEventResponse}
//	@Router		/admin/events/{id} [get]
func (h *EventHandler) GetEventByID(ctx echo.Context) error {
	eventID := ctx.Param("id")
	if eventID == "" {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	event, err := h.eventService.GetEventByID(ctx.Request().Context(), eventID)
	if err != nil {
		return err
	}

	eventResponse := dto.NewAdminEventResponse(event)

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", eventResponse, nil))
}

// CreateEvent
//
//	@Summary	Create Event
//	@Tags		[Admin] Event
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.CreateEventRequest	true	"Create Event Request"
//	@Success	200		{object}	response.Response{}
//	@Router		/admin/events [post]
func (h *EventHandler) CreateEvent(ctx echo.Context) error {
	var req dto.CreateEventRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	err := h.eventService.CreateEvent(ctx.Request().Context(), req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", nil, nil))
}

// UpdateEvent
//
//	@Summary	Update Event
//	@Tags		[Admin] Event
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string					true	"The Event ID"
//	@Param		request	body		dto.UpdateEventRequest	true	"Update Event Request"
//	@Success	200		{object}	response.Response{}
//	@Router		/admin/events/{id} [patch]
func (h *EventHandler) UpdateEvent(ctx echo.Context) error {
	var req dto.UpdateEventRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	err := h.eventService.UpdateEvent(ctx.Request().Context(), req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", nil, nil))
}

/* -------------------------------------------------------------------------- */
/*                                User Handler                                */
/* -------------------------------------------------------------------------- */

// GetUserEvents
//
//	@Summary	Get User Events
//	@Tags		[User] Event
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.Response{data=[]dto.EventListResponse}
//	@Router		/events [get]
func (h *EventHandler) GetUserEvents(ctx echo.Context) error {
	events, meta, err := h.eventService.GetUserEvents(ctx.Request().Context(), ctx.QueryParams())

	if err != nil {
		return err
	}

	eventResponse := dto.NewEventListResponse(events)

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", eventResponse, meta))
}

// GetUserEventByID
//
//	@Summary	Get Event By ID
//	@Tags		[User] Event
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"The Event ID"
//	@Success	200	{object}	response.Response{data=dto.EventResponse}
//	@Router		/events/{id} [get]
func (h *EventHandler) GetUserEventByID(ctx echo.Context) error {
	eventID := ctx.Param("id")
	if eventID == "" {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	event, err := h.eventService.GetUserEventByID(ctx.Request().Context(), eventID)
	if err != nil {
		return err
	}

	eventResponse := dto.NewEventResponse(event)

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", eventResponse, nil))
}

// RegisterEvent
//
//	@Summary	Register Event
//	@Tags		[User] Event
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.CreateEventRequest	true	"Create Event Request"
//	@Success	200		{object}	response.Response{}
//	@Router		/register-event [post]
func (h *EventHandler) RegisterEvent(ctx echo.Context) error {
	userID := ctx.Get("user_id").(string)
	var req dto.CreateEventRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	err := h.eventService.RegisterEvent(ctx.Request().Context(), req, userID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", nil, nil))
}
