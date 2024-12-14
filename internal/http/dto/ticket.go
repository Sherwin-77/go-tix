package dto

import (
	"github.com/google/uuid"
	"github.com/sherwin-77/go-tix/internal/entity"
)

type TicketResponse struct {
	ID       uuid.UUID `json:"id"`
	Category string    `json:"category"`
	Price    float64   `json:"price"`
	EventID  uuid.UUID `json:"event_id"`
}

func NewTicketResponse(ticket *entity.Ticket) *TicketResponse {
	return &TicketResponse{
		ID:       ticket.ID,
		Category: ticket.Category,
		Price:    ticket.Price,
		EventID:  ticket.EventID,
	}
}
