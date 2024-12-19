package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type EventApprovalMetadata struct {
}

type EventApproval struct {
	BaseEntity
	UserID      uuid.UUID                                 `json:"user_id" gorm:"type:uuid;not null"`
	Status      string                                    `json:"status" gorm:"type:varchar(20);not null"`
	Title       string                                    `json:"title" gorm:"type:varchar(255);not null"`
	Description string                                    `json:"description" gorm:"type:varchar(2047)"`
	Organizer   string                                    `json:"organizer" gorm:"type:varchar(255);not null"`
	Location    string                                    `json:"location" gorm:"type:varchar(2047);not null"`
	Longitude   float64                                   `json:"longitude" gorm:"type:decimal(8,6)"`
	Latitude    float64                                   `json:"latitude" gorm:"type:decimal(9,6)"`
	StartAt     time.Time                                 `json:"start_at" gorm:"type:timestamp(6) with time zone;not null"`
	EndAt       time.Time                                 `json:"end_at" gorm:"type:timestamp(6) with time zone;not null"`
	CreatedAt   time.Time                                 `json:"created_at" gorm:"type:timestamp(6) with time zone;not null"`
	Metadata    datatypes.JSONType[EventApprovalMetadata] `json:"metadata" gorm:"type:jsonb"`
	User        *User                                     `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
