package repository

import (
	"context"
	"github.com/sherwin-77/go-tix/internal/entity"
	"gorm.io/gorm"
)

type SaleInvoiceItemRepository interface {
	BaseRepository
	GetSaleInvoiceItems(ctx context.Context, tx *gorm.DB) ([]entity.SaleInvoiceItem, error)
	GetSaleInvoiceItemByID(ctx context.Context, tx *gorm.DB, id string) (*entity.SaleInvoiceItem, error)
	CreateSaleInvoiceItem(ctx context.Context, tx *gorm.DB, saleInvoiceItem *entity.SaleInvoiceItem) error
	UpdateSaleInvoiceItem(ctx context.Context, tx *gorm.DB, saleInvoiceItem *entity.SaleInvoiceItem) error
	DeleteSaleInvoiceItem(ctx context.Context, tx *gorm.DB, saleInvoiceItem *entity.SaleInvoiceItem) error
}

type saleInvoiceItemRepository struct {
	baseRepository
}

func NewSaleInvoiceItemRepository(db *gorm.DB) SaleInvoiceItemRepository {
	return &saleInvoiceItemRepository{baseRepository: baseRepository{db}}
}

func (r *saleInvoiceItemRepository) GetSaleInvoiceItems(ctx context.Context, tx *gorm.DB) ([]entity.SaleInvoiceItem, error) {
	var saleInvoiceItems []entity.SaleInvoiceItem
	if err := tx.WithContext(ctx).Find(&saleInvoiceItems).Error; err != nil {
		return nil, err
	}

	return saleInvoiceItems, nil
}

func (r *saleInvoiceItemRepository) GetSaleInvoiceItemByID(ctx context.Context, tx *gorm.DB, id string) (*entity.SaleInvoiceItem, error) {
	var saleInvoiceItem entity.SaleInvoiceItem
	if err := tx.WithContext(ctx).Where("id = ?", id).First(&saleInvoiceItem).Error; err != nil {
		return nil, err
	}

	return &saleInvoiceItem, nil
}

func (r *saleInvoiceItemRepository) CreateSaleInvoiceItem(ctx context.Context, tx *gorm.DB, saleInvoiceItem *entity.SaleInvoiceItem) error {
	if err := tx.WithContext(ctx).Create(saleInvoiceItem).Error; err != nil {
		return err
	}

	return nil
}

func (r *saleInvoiceItemRepository) UpdateSaleInvoiceItem(ctx context.Context, tx *gorm.DB, saleInvoiceItem *entity.SaleInvoiceItem) error {
	if err := tx.WithContext(ctx).Save(saleInvoiceItem).Error; err != nil {
		return err
	}

	return nil
}

func (r *saleInvoiceItemRepository) DeleteSaleInvoiceItem(ctx context.Context, tx *gorm.DB, saleInvoiceItem *entity.SaleInvoiceItem) error {
	if err := tx.WithContext(ctx).Delete(saleInvoiceItem).Error; err != nil {
		return err
	}

	return nil
}
