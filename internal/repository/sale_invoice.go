package repository

import (
	"context"
	"github.com/sherwin-77/go-tix/internal/entity"
	"gorm.io/gorm"
)

type SaleInvoiceRepository interface {
	BaseRepository
	SaleInvoiceNumberExists(ctx context.Context, tx *gorm.DB, number string) (bool, error)
	GetSaleInvoices(ctx context.Context, tx *gorm.DB) ([]entity.SaleInvoice, error)
	GetUserSaleInvoices(ctx context.Context, tx *gorm.DB, userID string) ([]entity.SaleInvoice, error)
	GetSaleInvoiceByID(ctx context.Context, tx *gorm.DB, id string) (*entity.SaleInvoice, error)
	CreateSaleInvoice(ctx context.Context, tx *gorm.DB, saleInvoice *entity.SaleInvoice) error
	UpdateSaleInvoice(ctx context.Context, tx *gorm.DB, saleInvoice *entity.SaleInvoice) error
	DeleteSaleInvoice(ctx context.Context, tx *gorm.DB, saleInvoice *entity.SaleInvoice) error
}

type saleInvoiceRepository struct {
	baseRepository
}

func NewSaleInvoiceRepository(db *gorm.DB) SaleInvoiceRepository {
	return &saleInvoiceRepository{baseRepository: baseRepository{db}}
}

func (r *saleInvoiceRepository) SaleInvoiceNumberExists(ctx context.Context, tx *gorm.DB, number string) (bool, error) {
	var count int64

	err := tx.WithContext(ctx).Model(&entity.SaleInvoice{}).Where("number = ?", number).Limit(1).Count(&count).Error

	return count > 0, err
}

func (r *saleInvoiceRepository) GetSaleInvoices(ctx context.Context, tx *gorm.DB) ([]entity.SaleInvoice, error) {
	var saleInvoices []entity.SaleInvoice

	if err := tx.WithContext(ctx).Find(&saleInvoices).Error; err != nil {
		return nil, err
	}

	return saleInvoices, nil
}

func (r *saleInvoiceRepository) GetUserSaleInvoices(ctx context.Context, tx *gorm.DB, userID string) ([]entity.SaleInvoice, error) {
	var saleInvoices []entity.SaleInvoice

	if err := tx.WithContext(ctx).Where("user_id = ?", userID).Find(&saleInvoices).Error; err != nil {
		return nil, err
	}

	return saleInvoices, nil
}

func (r *saleInvoiceRepository) GetSaleInvoiceByID(ctx context.Context, tx *gorm.DB, id string) (*entity.SaleInvoice, error) {
	var saleInvoice entity.SaleInvoice

	if err := tx.WithContext(ctx).Where("id = ?", id).First(&saleInvoice).Error; err != nil {
		return nil, err
	}

	return &saleInvoice, nil
}

func (r *saleInvoiceRepository) CreateSaleInvoice(ctx context.Context, tx *gorm.DB, saleInvoice *entity.SaleInvoice) error {
	if err := tx.WithContext(ctx).Create(saleInvoice).Error; err != nil {
		return err
	}

	return nil
}

func (r *saleInvoiceRepository) UpdateSaleInvoice(ctx context.Context, tx *gorm.DB, saleInvoice *entity.SaleInvoice) error {
	if err := tx.WithContext(ctx).Save(saleInvoice).Error; err != nil {
		return err
	}

	return nil
}

func (r *saleInvoiceRepository) DeleteSaleInvoice(ctx context.Context, tx *gorm.DB, saleInvoice *entity.SaleInvoice) error {
	if err := tx.WithContext(ctx).Delete(saleInvoice).Error; err != nil {
		return err
	}

	return nil
}
