package entity

import (
	"github.com/google/uuid"
)

type EventApprovalTicketsMetadata struct {
}

type EventApprovalTicket struct {
	BaseEntity
	EventApprovalID uuid.UUID `json:"event_approval_id" gorm:"type:uuid;not null"`
	Category        string    `json:"category" gorm:"type:varchar(100)"`
	Price           float64   `json:"price" gorm:"type:decimal(12,2);not null;default:0"`

	EventApproval *EventApproval `json:"event_approval,omitempty" gorm:"foreignKey:EventApprovalID"`
}
