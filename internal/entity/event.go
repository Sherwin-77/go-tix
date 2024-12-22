package entity

import (
	"github.com/guregu/null/v5"
	"gorm.io/datatypes"
)

type Event struct {
	BaseEntity
	Status      string         `json:"status" gorm:"type:varchar(20);not null"`
	Title       string         `json:"title" gorm:"type:varchar(255);not null"`
	Description null.String    `json:"description" gorm:"type:varchar(2047)"`
	Organizer   string         `json:"organizer" gorm:"type:varchar(255);not null"`
	Location    null.String    `json:"location" gorm:"type:varchar(2047)"`
	Longitude   null.Float     `json:"longitude" gorm:"type:decimal(9,6)"`
	Latitude    null.Float     `json:"latitude" gorm:"type:decimal(8,6)"`
	StartAt     datatypes.Date `json:"start_at" gorm:"type:date;not null"`
	EndAt       datatypes.Date `json:"end_at" gorm:"type:date;not null"`

	Tickets []*Ticket `json:"tickets,omitempty" gorm:"foreignKey:EventID"`
}

type EventWithMinMaxPrice struct {
	Event
	MinPrice float64 `json:"min_price"`
	MaxPrice float64 `json:"max_price"`
}

func (e *Event) TableName() string {
	return "events"
}
