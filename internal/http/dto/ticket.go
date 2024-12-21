package dto

import (
	"github.com/sherwin-77/go-tix/internal/entity"
)

type TicketResponse struct {
	ID       string  `json:"id"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

func NewTicketResponse(ticket *entity.Ticket) *TicketResponse {
	return &TicketResponse{
		ID:       ticket.ID.String(),
		Category: ticket.Category,
		Price:    ticket.Price,
	}
}
