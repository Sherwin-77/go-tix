package repository

import (
	"context"
	"github.com/sherwin-77/go-tix/internal/entity"
	"gorm.io/gorm"
)

type EventRepository interface {
	BaseRepository
	GetEvents(ctx context.Context, tx *gorm.DB) ([]entity.Event, error)
	GetEventsWithMinMaxPrice(ctx context.Context, tx *gorm.DB) ([]entity.EventWithMinMaxPrice, error)
	GetEventByID(ctx context.Context, tx *gorm.DB, id string) (*entity.Event, error)
	CreateEvent(ctx context.Context, tx *gorm.DB, event *entity.Event) error
	UpdateEvent(ctx context.Context, tx *gorm.DB, event *entity.Event) error
	DeleteEvent(ctx context.Context, tx *gorm.DB, event *entity.Event) error
}

type eventRepository struct {
	baseRepository
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{baseRepository: baseRepository{db}}
}

func (r *eventRepository) GetEvents(ctx context.Context, tx *gorm.DB) ([]entity.Event, error) {
	var events []entity.Event

	if err := tx.WithContext(ctx).Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

func (r *eventRepository) GetEventsWithMinMaxPrice(ctx context.Context, tx *gorm.DB) ([]entity.EventWithMinMaxPrice, error) {
	var events []entity.EventWithMinMaxPrice

	if err := tx.WithContext(ctx).Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

func (r *eventRepository) GetEventByID(ctx context.Context, tx *gorm.DB, id string) (*entity.Event, error) {
	var event entity.Event

	if err := tx.WithContext(ctx).Where("id = ?", id).First(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *eventRepository) CreateEvent(ctx context.Context, tx *gorm.DB, event *entity.Event) error {
	if err := tx.WithContext(ctx).Create(event).Error; err != nil {
		return err
	}

	return nil
}

func (r *eventRepository) UpdateEvent(ctx context.Context, tx *gorm.DB, event *entity.Event) error {
	if err := tx.WithContext(ctx).Save(event).Error; err != nil {
		return err
	}

	return nil
}

func (r *eventRepository) DeleteEvent(ctx context.Context, tx *gorm.DB, event *entity.Event) error {
	if err := tx.WithContext(ctx).Delete(event).Error; err != nil {
		return err
	}

	return nil
}
