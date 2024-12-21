package dto

import "github.com/sherwin-77/go-tix/internal/entity"

type EventApprovalTicketResponse struct {
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

func NewEventApprovalTicketResponse(ticket *entity.EventApprovalTicket) *EventApprovalTicketResponse {
	return &EventApprovalTicketResponse{
		Category: ticket.Category,
		Price:    ticket.Price,
	}
}
