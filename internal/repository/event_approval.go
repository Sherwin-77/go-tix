package repository

import (
	"context"

	"github.com/sherwin-77/go-tix/internal/entity"
	"gorm.io/gorm"
)

type EventApprovalRepository interface {
	BaseRepository
	GetEventApprovals(ctx context.Context) ([]entity.EventApproval, error)
	GetEventApprovalByID(ctx context.Context, id string) (*entity.EventApproval, error)
	HandleEventApproval(ctx context.Context, eventApproval *entity.EventApproval) error
	CreateEventApproval(ctx context.Context, eventApproval *entity.EventApproval) error
	UpdateEventApproval(ctx context.Context, eventApproval *entity.EventApproval) error
	DeleteEventApproval(ctx context.Context, id string) error
}

type eventApprovalRepository struct {
	baseRepository
}

func NewEventApprovalRepository(db *gorm.DB) EventApprovalRepository {
	return &eventApprovalRepository{
		baseRepository: baseRepository{
			db: db,
		},
	}
}

// GetEventApprovals retrieves all event approvals
func (r *eventApprovalRepository) GetEventApprovals(ctx context.Context) ([]entity.EventApproval, error) {
	var eventApprovals []entity.EventApproval
	return eventApprovals, r.db.WithContext(ctx).Find(&eventApprovals).Error
}

// GetEventApprovalByID retrieves an event approval by ID
func (r *eventApprovalRepository) GetEventApprovalByID(ctx context.Context, id string) (*entity.EventApproval, error) {
	var approval entity.EventApproval
	return &approval, r.db.WithContext(ctx).First(&approval, "id = ?", id).Error

}

// HandleEventApproval updates the status of an event approval
func (r *eventApprovalRepository) HandleEventApproval(ctx context.Context, eventApproval *entity.EventApproval) error {
	return r.db.WithContext(ctx).Save(eventApproval).Error
}

// CreateEventApproval membuat event approval baru
func (r *eventApprovalRepository) CreateEventApproval(ctx context.Context, eventApproval *entity.EventApproval) error {
	return r.db.WithContext(ctx).Create(eventApproval).Error
}

// UpdateEventApproval updates an event approval
func (r *eventApprovalRepository) UpdateEventApproval(ctx context.Context, updatedEventApproval *entity.EventApproval) error {
	return r.db.WithContext(ctx).Save(updatedEventApproval).Error
}

// DeleteEventApproval deletes an event approval
func (r *eventApprovalRepository) DeleteEventApproval(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&entity.EventApproval{}, "id = ?", id).Error
}
