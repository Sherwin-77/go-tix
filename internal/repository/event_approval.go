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
	return &eventApprovalRepository{baseRepository: baseRepository{db}}
}

// GetEventApprovals mengambil semua event approval
func (r *eventApprovalRepository) GetEventApprovals(ctx context.Context) ([]entity.EventApproval, error) {
	var eventApprovals []entity.EventApproval
	err := r.db.WithContext(ctx).Find(&eventApprovals).Error
	if err != nil {
		return nil, err
	}
	return eventApprovals, nil
}

// GetEventApprovalByID mengambil detail event approval berdasarkan ID
func (r *eventApprovalRepository) GetEventApprovalByID(ctx context.Context, id string) (*entity.EventApproval, error) {
	var eventApproval entity.EventApproval
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&eventApproval).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &eventApproval, nil
}

// HandleEventApproval menangani persetujuan atau penolakan event
func (r *eventApprovalRepository) HandleEventApproval(ctx context.Context, eventApproval *entity.EventApproval) error {
	// Update status event approval
	err := r.db.WithContext(ctx).Save(eventApproval).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateEventApproval membuat event approval baru
func (r *eventApprovalRepository) CreateEventApproval(ctx context.Context, eventApproval *entity.EventApproval) error {
	err := r.db.WithContext(ctx).Create(eventApproval).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateEventApproval memperbarui event approval
func (r *eventApprovalRepository) UpdateEventApproval(ctx context.Context, eventApproval *entity.EventApproval) error {
	err := r.db.WithContext(ctx).Save(eventApproval).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteEventApproval menghapus event approval
func (r *eventApprovalRepository) DeleteEventApproval(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.EventApproval{}).Error
	if err != nil {
		return err
	}
	return nil
}
