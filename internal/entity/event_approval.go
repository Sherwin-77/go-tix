package entity

import (
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"gorm.io/datatypes"
)

type EventApprovalMetadata struct {
}

type EventApproval struct {
	BaseEntity
	UserID      uuid.UUID                                 `json:"user_id" gorm:"type:uuid;not null"`
	Status      string                                    `json:"status" gorm:"type:varchar(20);not null"`
	Title       string                                    `json:"title" gorm:"type:varchar(255);not null"`
	Description null.String                               `json:"description" gorm:"type:varchar(2047)"`
	Organizer   string                                    `json:"organizer" gorm:"type:varchar(255);not null"`
	Location    null.String                               `json:"location" gorm:"type:varchar(2047)"`
	Longitude   null.Float                                `json:"longitude" gorm:"type:decimal(8,6)"`
	Latitude    null.Float                                `json:"latitude" gorm:"type:decimal(9,6)"`
	StartAt     datatypes.Date                            `json:"start_at" gorm:"type:timestamp(6) with time zone;not null"`
	EndAt       datatypes.Date                            `json:"end_at" gorm:"type:timestamp(6) with time zone;not null"`
	Metadata    datatypes.JSONType[EventApprovalMetadata] `json:"metadata" gorm:"type:jsonb"`

	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`

	EventApprovalTickets []*EventApprovalTicket `json:"event_approvals_tickets,omitempty" gorm:"foreignKey:EventApprovalID"`
}
