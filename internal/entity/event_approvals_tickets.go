package entity

import (
	"time"

	"github.com/google/uuid"
)

type EventApprovalTicketsMetadata struct {
	AdditionalInfo string `json:"additional_info"`
}

type EventApprovalTicket struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	EventApprovalID uuid.UUID `json:"event_approval_id" gorm:"type:uuid;not null"`
	Category        string    `json:"category" gorm:"type:varchar(100)"`
	Price           float64   `json:"price" gorm:"type:decimal(12,2);not null;default:0"`
	CreatedAt       time.Time `json:"created_at" gorm:"type:timestamp(6) with time zone"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"type:timestamp(6) with time zone"`

	EventApproval *EventApproval `json:"event_approval,omitempty" gorm:"foreignKey:EventApprovalID"`
}
