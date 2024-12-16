package repository

import (
	"context"

	"github.com/sherwin-77/go-tix/internal/entity"
	"gorm.io/gorm"
)

type EventApprovalTicketRepository interface {
	BaseRepository
	GetEventApprovalsTickets(ctx context.Context) ([]entity.EventApproval, error)
	GetEventApprovalsByEventIDTickets(ctx context.Context, id int64) ([]entity.EventApproval, error)
	CreateEventApprovalTickets(ctx context.Context, eventApproval *entity.EventApproval) error
	UpdateEventApprovalTickets(ctx context.Context, eventApproval *entity.EventApproval) error
	DeleteEventApprovalTickets(ctx context.Context, id int64) error
}

type eventApprovalTicketRepository struct {
	baseRepository
}

func NewEventApprovalTicketRepository(db *gorm.DB) EventApprovalTicketRepository {
	return &eventApprovalTicketRepository{baseRepository{db}}
}

func (r *eventApprovalTicketRepository) GetEventApprovalsTickets(ctx context.Context) ([]entity.EventApproval, error) {
	var eventApprovals []entity.EventApproval
	err := r.db.WithContext(ctx).Find(&eventApprovals).Error
	if err != nil {
		return nil, err
	}

	return eventApprovals, nil
}

func (r *eventApprovalTicketRepository) GetEventApprovalsByEventIDTickets(ctx context.Context, id int64) ([]entity.EventApproval, error) {
	var eventApprovals []entity.EventApproval
	err := r.db.WithContext(ctx).Where("event_id = ?", id).Find(&eventApprovals).Error
	if err != nil {
		return nil, err
	}

	return eventApprovals, nil
}

func (r *eventApprovalTicketRepository) CreateEventApprovalTickets(ctx context.Context, eventApproval *entity.EventApproval) error {
	err := r.db.WithContext(ctx).Create(eventApproval).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *eventApprovalTicketRepository) UpdateEventApprovalTickets(ctx context.Context, eventApproval *entity.EventApproval) error {
	err := r.db.WithContext(ctx).Save(eventApproval).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *eventApprovalTicketRepository) DeleteEventApprovalTickets(ctx context.Context, id int64) error {
	err := r.db.WithContext(ctx).Delete(&entity.EventApproval{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
