package repository

import (
	"context"
	"errors"

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

	if err := r.db.WithContext(ctx).Find(&eventApprovals).Error; err != nil {
		return nil, err
	}

	return eventApprovals, nil
}

// GetEventApprovalByID retrieves an event approval by ID
func (r *eventApprovalRepository) GetEventApprovalByID(ctx context.Context, id string) (*entity.EventApproval, error) {
	var approval entity.EventApproval

	err := r.db.WithContext(ctx).Where("id = ?", id).First(&approval).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &approval, nil
}

// HandleEventApproval updates the status of an event approval
func (r *eventApprovalRepository) HandleEventApproval(ctx context.Context, eventApproval *entity.EventApproval) error {
	return r.db.WithContext(ctx).Save(eventApproval).Error
}

// CreateEventApproval membuat event approval baru
func (r *eventApprovalRepository) CreateEventApproval(ctx context.Context, eventApproval *entity.EventApproval) error {
	err := r.db.WithContext(ctx).Create(eventApproval).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateEventApproval updates an event approval
func (r *eventApprovalRepository) UpdateEventApproval(ctx context.Context, updatedEventApproval *entity.EventApproval) error {
	return r.db.WithContext(ctx).Save(updatedEventApproval).Error
}

// DeleteEventApproval deletes an event approval
func (r *eventApprovalRepository) DeleteEventApproval(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Delete(&entity.EventApproval{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
