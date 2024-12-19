package enum

type EventStatus string

const (
	EventStatusDisabled EventStatus = "disabled"
	EventStatusActive   EventStatus = "active"
	EventStatusFinished EventStatus = "finished"
)
