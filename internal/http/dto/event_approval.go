package dto

import (
	"github.com/guregu/null/v5"
	"github.com/sherwin-77/go-tix/internal/entity"
	"gorm.io/datatypes"
)

type HandleEventApprovalRequest struct {
	ID     string `param:"id" validate:"required,uuid" swaggerignore:"true"`
	Action string `json:"action" validate:"required,oneof=approve reject"`
}

type EventApprovalResponse struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Title     string `json:"title"`
	Organizer string `json:"organizer"`
	Location  null.String
	Longitude null.Float
	Latitude  null.Float
	StartAt   datatypes.Date
	EndAt     datatypes.Date

	User                 *UserResponse                  `json:"user,omitempty"`
	EventApprovalTickets []*EventApprovalTicketResponse `json:"event_approvals_tickets,omitempty"`
}

func NewEventApprovalResponse(eventApproval *entity.EventApproval) EventApprovalResponse {

	response := EventApprovalResponse{
		ID:        eventApproval.ID.String(),
		Status:    eventApproval.Status,
		Title:     eventApproval.Title,
		Organizer: eventApproval.Organizer,
		Location:  eventApproval.Location,
		Longitude: eventApproval.Longitude,
		Latitude:  eventApproval.Latitude,
		StartAt:   eventApproval.StartAt,
		EndAt:     eventApproval.EndAt,
	}

	if eventApproval.User != nil {
		user := NewUserResponse(eventApproval.User)
		response.User = &user
	}

	for _, ticket := range eventApproval.EventApprovalTickets {
		ticketResponse := NewEventApprovalTicketResponse(ticket)
		response.EventApprovalTickets = append(response.EventApprovalTickets, ticketResponse)
	}

	return response
}

func NewEventApprovalsResponse(eventApprovals []entity.EventApproval) []EventApprovalResponse {
	res := make([]EventApprovalResponse, 0)
	for _, eventApproval := range eventApprovals {
		res = append(res, NewEventApprovalResponse(&eventApproval))
	}
	return res
}
