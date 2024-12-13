package repository

import (
	"context"

	"github.com/sherwin-77/go-tix/internal/entity"
	"gorm.io/gorm"
)

type EventApprovalRepository interface {
	BaseRepository
	GetEventApprovals(ctx context.Context, tx *gorm.DB) ([]entity.EventApproval, error)
	GetEventApprovalsFiltered(ctx context.Context, tx *gorm.DB, limit int, offset int, order interface{}, query interface{}, args ...interface{}) ([]entity.EventApproval, error)
	GetEventApprovalByID(ctx context.Context, tx *gorm.DB, id string) (*entity.EventApproval, error)
	CreateEventApproval(ctx context.Context, tx *gorm.DB, eventApproval *entity.EventApproval) error
	UpdateEventApproval(ctx context.Context, tx *gorm.DB, eventApproval *entity.EventApproval) error
	DeleteEventApproval(ctx context.Context, tx *gorm.DB, eventApproval *entity.EventApproval) error
}

type eventApprovalRepository struct {
	baseRepository
}

func NewEventApprovalRepository(db *gorm.DB) EventApprovalRepository {
	return &eventApprovalRepository{baseRepository{db}}
}

func (r *eventApprovalRepository) GetEventApprovals(ctx context.Context, tx *gorm.DB) ([]entity.EventApproval, error) {
	var eventApprovals []entity.EventApproval

	if err := tx.WithContext(ctx).Find(&eventApprovals).Error; err != nil {
		return nil, err
	}

	return eventApprovals, nil
}

func (r *eventApprovalRepository) GetEventApprovalsFiltered(ctx context.Context, tx *gorm.DB, limit int, offset int, order interface{}, query interface{}, args ...interface{}) ([]entity.EventApproval, error) {
	var eventApprovals []entity.EventApproval

	if err := tx.WithContext(ctx).Where(query, args...).Limit(limit).Offset(offset).Order(order).Find(&eventApprovals).Error; err != nil {
		return nil, err
	}

	return eventApprovals, nil
}

func (r *eventApprovalRepository) GetEventApprovalByID(ctx context.Context, tx *gorm.DB, id string) (*entity.EventApproval, error) {
	var eventApproval entity.EventApproval

	if err := tx.WithContext(ctx).Where("id = ?", id).First(&eventApproval).Error; err != nil {
		return nil, err
	}

	return &eventApproval, nil
}

func (r *eventApprovalRepository) CreateEventApproval(ctx context.Context, tx *gorm.DB, eventApproval *entity.EventApproval) error {
	if err := tx.WithContext(ctx).Create(eventApproval).Error; err != nil {
		return err
	}

	return nil
}

func (r *eventApprovalRepository) UpdateEventApproval(ctx context.Context, tx *gorm.DB, eventApproval *entity.EventApproval) error {
	if err := tx.WithContext(ctx).Save(eventApproval).Error; err != nil {
		return err
	}

	return nil
}

func (r *eventApprovalRepository) DeleteEventApproval(ctx context.Context, tx *gorm.DB, eventApproval *entity.EventApproval) error {
	if err := tx.WithContext(ctx).Delete(eventApproval).Error; err != nil {
		return err
	}

	return nil
}
