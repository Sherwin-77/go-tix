package entity

import (
	"gorm.io/datatypes"
)

type Event struct {
	BaseEntity
	Title       string         `json:"title" gorm:"type:varchar(255);not null"`
	Description string         `json:"description" gorm:"type:varchar(2047)"`
	Organizer   string         `json:"organizer" gorm:"type:varchar(255);not null"`
	Location    string         `json:"location" gorm:"type:varchar(2047)"`
	Longitude   float64        `json:"longitude" gorm:"type:decimal(8,6)"`
	Latitude    float64        `json:"latitude" gorm:"type:decimal(9,6)"`
	StartAt     datatypes.Date `json:"start_at" gorm:"type:date;not null"`
	EndAt       datatypes.Date `json:"end_at" gorm:"type:date;not null"`

	Tickets []*Ticket `json:"tickets,omitempty" gorm:"foreignKey:EventID"`
}
