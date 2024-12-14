package dto

import (
	"github.com/sherwin-77/go-tix/internal/entity"
	"gorm.io/datatypes"
)

type EventListResponse struct {
	ID          string         `json:"id"`
	Description string         `json:"description"`
	Title       string         `json:"title"`
	Organizer   string         `json:"organizer"`
	Location    string         `json:"location"`
	Longitude   float64        `json:"longitude"`
	Latitude    float64        `json:"latitude"`
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
	Description string            `json:"description"`
	Title       string            `json:"title"`
	Organizer   string            `json:"organizer"`
	Location    string            `json:"location"`
	Longitude   float64           `json:"longitude"`
	Latitude    float64           `json:"latitude"`
	StartAt     datatypes.Date    `json:"start_at"`
	EndAt       datatypes.Date    `json:"end_at"`
	Tickets     []*TicketResponse `json:"tickets,omitempty"`
}

type AdminEventResponse struct {
	EventResponse
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewEventResponse(event entity.Event) EventResponse {
	response := EventResponse{
		ID:          event.ID.String(),
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

func NewAdminEventResponse(event entity.Event) AdminEventResponse {
	return AdminEventResponse{
		EventResponse: NewEventResponse(event),
		CreatedAt:     event.CreatedAt.String(),
		UpdatedAt:     event.UpdatedAt.String(),
	}
}

func NewEventListResponse(events []entity.EventWithMinMaxPrice) []EventListResponse {
	var res []EventListResponse
	for _, event := range events {
		res = append(res, EventListResponse{
			ID:          event.ID.String(),
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
	var res []AdminEventListResponse
	for _, event := range events {
		res = append(res, AdminEventListResponse{
			EventListResponse: EventListResponse{
				ID:          event.ID.String(),
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
