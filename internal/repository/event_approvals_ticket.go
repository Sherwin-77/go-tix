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
	return eventApprovals, r.db.WithContext(ctx).Find(&eventApprovals).Error
}

func (r *eventApprovalTicketRepository) GetEventApprovalsByEventIDTickets(ctx context.Context, id int64) ([]entity.EventApproval, error) {
	var eventApprovals []entity.EventApproval
	return eventApprovals, r.db.WithContext(ctx).Where("event_id = ?", id).Find(&eventApprovals).Error
}

func (r *eventApprovalTicketRepository) CreateEventApprovalTickets(ctx context.Context, eventApproval *entity.EventApproval) error {
	return r.db.WithContext(ctx).Create(eventApproval).Error

}

func (r *eventApprovalTicketRepository) UpdateEventApprovalTickets(ctx context.Context, eventApproval *entity.EventApproval) error {
	return r.db.WithContext(ctx).Save(eventApproval).Error
}

func (r *eventApprovalTicketRepository) DeleteEventApprovalTickets(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.EventApproval{}, id).Error
}
