package repository

import (
	"context"
	"github.com/sherwin-77/go-tix/internal/entity"
	"gorm.io/gorm"
)

type SnapPaymentRepository interface {
	BaseRepository
	ExternalIDExists(ctx context.Context, tx *gorm.DB, externalID string) (bool, error)
	GetByExternalID(ctx context.Context, tx *gorm.DB, externalID string) (*entity.SnapPayment, error)
	CreateSnapPayment(ctx context.Context, tx *gorm.DB, snapPayment *entity.SnapPayment) error
	UpdateSnapPayment(ctx context.Context, tx *gorm.DB, snapPayment *entity.SnapPayment) error
}

type snapPaymentRepository struct {
	baseRepository
}

func NewSnapPaymentRepository(db *gorm.DB) SnapPaymentRepository {
	return &snapPaymentRepository{baseRepository: baseRepository{db}}
}

func (r *snapPaymentRepository) ExternalIDExists(ctx context.Context, tx *gorm.DB, externalID string) (bool, error) {
	var count int64

	err := tx.WithContext(ctx).Model(&entity.SnapPayment{}).Where("external_id = ?", externalID).Limit(1).Count(&count).Error

	return count > 0, err
}

func (r *snapPaymentRepository) GetByExternalID(ctx context.Context, tx *gorm.DB, externalID string) (*entity.SnapPayment, error) {
	var snapPayment entity.SnapPayment

	if err := tx.WithContext(ctx).Where("external_id = ?", externalID).First(&snapPayment).Error; err != nil {
		return nil, err
	}

	return &snapPayment, nil
}

func (r *snapPaymentRepository) CreateSnapPayment(ctx context.Context, tx *gorm.DB, snapPayment *entity.SnapPayment) error {
	if err := tx.WithContext(ctx).Create(snapPayment).Error; err != nil {
		return err
	}

	return nil
}

func (r *snapPaymentRepository) UpdateSnapPayment(ctx context.Context, tx *gorm.DB, snapPayment *entity.SnapPayment) error {
	if err := tx.WithContext(ctx).Save(snapPayment).Error; err != nil {
		return err
	}

	return nil
}
