package dto

import (
	"github.com/guregu/null/v5"
	"github.com/sherwin-77/go-tix/internal/entity"
	"gorm.io/datatypes"
)

/* -------------------------------------------------------------------------- */
/*                                   Request                                  */
/* -------------------------------------------------------------------------- */

type CreateEventRequest struct {
	Title       string              `json:"title" validate:"required,lte=255"`
	Description *string             `json:"description,omitempty" validate:"omitnil,lte=2047"`
	Organizer   string              `json:"organizer" validate:"required,lte=255"`
	Location    *string             `json:"location" validate:"omitnil,lte=2047"`
	Longitude   *float64            `json:"longitude" validate:"omitnil,longitude"`
	Latitude    *float64            `json:"latitude" validate:"omitnil,latitude"`
	StartAt     datatypes.Date      `json:"start_at" validate:"required"`
	EndAt       datatypes.Date      `json:"end_at" validate:"required"`
	Tickets     []TicketRequestItem `json:"tickets"`
}

type TicketRequestItem struct {
	Category string   `json:"category"`
	Price    *float64 `json:"price" validate:"required,gte=0,numeric"`
}

type UpdateEventRequest struct {
	ID        string `param:"id" validate:"required,uuid" swaggerignore:"true"`
	IsEnabled *bool  `json:"is_enabled" validate:"required,boolean"`
}

/* -------------------------------------------------------------------------- */
/*                                  Response                                  */
/* -------------------------------------------------------------------------- */

type EventListResponse struct {
	ID          string         `json:"id"`
	Status      string         `json:"status"`
	Description null.String    `json:"description"`
	Title       string         `json:"title"`
	Organizer   string         `json:"organizer"`
	Location    null.String    `json:"location"`
	Longitude   null.Float     `json:"longitude"`
	Latitude    null.Float     `json:"latitude"`
	StartAt     datatypes.Date `json:"start_at"`
	EndAt       datatypes.Date `json:"end_at"`
	MinPrice    float64        `json:"min_price"`
	MaxPrice    float64        `json:"max_price"`
}

type AdminEventListResponse struct {
	EventListResponse
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type EventResponse struct {
	ID          string            `json:"id"`
	Status      string            `json:"status"`
	Description null.String       `json:"description"`
	Title       string            `json:"title"`
	Organizer   string            `json:"organizer"`
	Location    null.String       `json:"location"`
	Longitude   null.Float        `json:"longitude"`
	Latitude    null.Float        `json:"latitude"`
	StartAt     datatypes.Date    `json:"start_at"`
	EndAt       datatypes.Date    `json:"end_at"`
	Tickets     []*TicketResponse `json:"tickets,omitempty"`
}

type AdminEventResponse struct {
	EventResponse
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewEventResponse(event *entity.Event) EventResponse {
	response := EventResponse{
		ID:          event.ID.String(),
		Status:      event.Status,
		Description: event.Description,
		Title:       event.Title,
		Organizer:   event.Organizer,
		Location:    event.Location,
		Longitude:   event.Longitude,
		Latitude:    event.Latitude,
		StartAt:     event.StartAt,
		EndAt:       event.EndAt,
	}
	for _, ticket := range event.Tickets {
		response.Tickets = append(response.Tickets, NewTicketResponse(ticket))
	}

	return response
}

func NewAdminEventResponse(event *entity.Event) AdminEventResponse {
	return AdminEventResponse{
		EventResponse: NewEventResponse(event),
		CreatedAt:     event.CreatedAt.String(),
		UpdatedAt:     event.UpdatedAt.String(),
	}
}

func NewEventListResponse(events []entity.EventWithMinMaxPrice) []EventListResponse {
	var res = make([]EventListResponse, 0)
	for _, event := range events {
		res = append(res, EventListResponse{
			ID:          event.ID.String(),
			Status:      event.Status,
			Description: event.Description,
			Title:       event.Title,
			Organizer:   event.Organizer,
			Location:    event.Location,
			Longitude:   event.Longitude,
			Latitude:    event.Latitude,
			StartAt:     event.StartAt,
			EndAt:       event.EndAt,
			MinPrice:    event.MinPrice,
			MaxPrice:    event.MaxPrice,
		})
	}
	return res
}

func NewAdminEventListResponse(events []entity.EventWithMinMaxPrice) []AdminEventListResponse {
	var res = make([]AdminEventListResponse, 0)
	for _, event := range events {
		res = append(res, AdminEventListResponse{
			EventListResponse: EventListResponse{
				ID:          event.ID.String(),
				Status:      event.Status,
				Description: event.Description,
				Title:       event.Title,
				Organizer:   event.Organizer,
				Location:    event.Location,
				Longitude:   event.Longitude,
				Latitude:    event.Latitude,
				StartAt:     event.StartAt,
				EndAt:       event.EndAt,
				MinPrice:    event.MinPrice,
				MaxPrice:    event.MaxPrice,
			},
			CreatedAt: event.CreatedAt.String(),
			UpdatedAt: event.UpdatedAt.String(),
		})
	}
	return res
}
