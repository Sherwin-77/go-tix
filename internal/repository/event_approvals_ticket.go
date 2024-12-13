package repository

import (
	"context"

	"github.com/sherwin-77/go-tix/internal/entity"
	"gorm.io/gorm"
)

type EventApprovalTicketRepository interface {
	BaseRepository
	GetEventApprovalTickets(ctx context.Context, tx *gorm.DB) ([]entity.EventApprovalTicket, error)
	GetEventApprovalTicketsFiltered(ctx context.Context, tx *gorm.DB, limit int, offset int, order interface{}, query interface{}, args ...interface{}) ([]entity.EventApprovalTicket, error)
	GetEventApprovalTicketByID(ctx context.Context, tx *gorm.DB, id string) (*entity.EventApprovalTicket, error)
	CreateEventApprovalTicket(ctx context.Context, tx *gorm.DB, eventApprovalTicket *entity.EventApprovalTicket) error
	UpdateEventApprovalTicket(ctx context.Context, tx *gorm.DB, eventApprovalTicket *entity.EventApprovalTicket) error
	DeleteEventApprovalTicket(ctx context.Context, tx *gorm.DB, eventApprovalTicket *entity.EventApprovalTicket) error
}

type eventApprovalTicketRepository struct {
	baseRepository
}

func NewEventApprovalTicketRepository(db *gorm.DB) EventApprovalTicketRepository {
	return &eventApprovalTicketRepository{baseRepository{db}}
}

func (r *eventApprovalTicketRepository) GetEventApprovalTickets(ctx context.Context, tx *gorm.DB) ([]entity.EventApprovalTicket, error) {
	var eventApprovalTickets []entity.EventApprovalTicket

	if err := tx.WithContext(ctx).Find(&eventApprovalTickets).Error; err != nil {
		return nil, err
	}

	return eventApprovalTickets, nil
}

func (r *eventApprovalTicketRepository) GetEventApprovalTicketsFiltered(ctx context.Context, tx *gorm.DB, limit int, offset int, order interface{}, query interface{}, args ...interface{}) ([]entity.EventApprovalTicket, error) {
	var eventApprovalTickets []entity.EventApprovalTicket

	if err := tx.WithContext(ctx).Where(query, args...).Limit(limit).Offset(offset).Order(order).Find(&eventApprovalTickets).Error; err != nil {
		return nil, err
	}

	return eventApprovalTickets, nil
}

func (r *eventApprovalTicketRepository) GetEventApprovalTicketByID(ctx context.Context, tx *gorm.DB, id string) (*entity.EventApprovalTicket, error) {
	var eventApprovalTicket entity.EventApprovalTicket

	if err := tx.WithContext(ctx).Where("id = ?", id).First(&eventApprovalTicket).Error; err != nil {
		return nil, err
	}

	return &eventApprovalTicket, nil
}

func (r *eventApprovalTicketRepository) CreateEventApprovalTicket(ctx context.Context, tx *gorm.DB, eventApprovalTicket *entity.EventApprovalTicket) error {
	if err := tx.WithContext(ctx).Create(eventApprovalTicket).Error; err != nil {
		return err
	}

	return nil
}

func (r *eventApprovalTicketRepository) UpdateEventApprovalTicket(ctx context.Context, tx *gorm.DB, eventApprovalTicket *entity.EventApprovalTicket) error {
	if err := tx.WithContext(ctx).Save(eventApprovalTicket).Error; err != nil {
		return err
	}

	return nil
}

func (r *eventApprovalTicketRepository) DeleteEventApprovalTicket(ctx context.Context, tx *gorm.DB, eventApprovalTicket *entity.EventApprovalTicket) error {
	if err := tx.WithContext(ctx).Delete(eventApprovalTicket).Error; err != nil {
		return err
	}

	return nil
}
