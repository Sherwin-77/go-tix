package service

import (
	"context"
	"fmt"
	"github.com/guregu/null/v5"
	"github.com/sherwin-77/go-tix/internal/entity"
	"github.com/sherwin-77/go-tix/internal/enum"
	"github.com/sherwin-77/go-tix/internal/repository"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, tx *gorm.DB, saleInvoice *entity.SaleInvoice) (*entity.SnapPayment, error)
}

type transactionService struct {
	midtransService       MidtransService
	snapPaymentRepository repository.SnapPaymentRepository
	randomizer            *rand.Rand
}

func NewTransactionService(midtransService MidtransService, snapPaymentRepository repository.SnapPaymentRepository) TransactionService {
	return &transactionService{
		midtransService,
		snapPaymentRepository,
		rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *transactionService) generateOrderID() string {
	return fmt.Sprintf("TIX-%s-%08d", time.Now().Format("20060102150405"), s.randomizer.Intn(99999999))
}

func (s *transactionService) CreateTransaction(ctx context.Context, tx *gorm.DB, saleInvoice *entity.SaleInvoice) (*entity.SnapPayment, error) {
	snapPayment := &entity.SnapPayment{
		UserID:        saleInvoice.UserID,
		SaleInvoiceID: saleInvoice.ID,
		ExternalID:    saleInvoice.Number,
		Amount:        saleInvoice.Total,
		Status:        string(enum.SnapPaymentStatusPending),
		Method:        "snap", // Notice hardcode method, change if needed
	}

	for {
		snapPayment.ExternalID = s.generateOrderID()
		exist, err := s.snapPaymentRepository.ExternalIDExists(ctx, tx, snapPayment.ExternalID)
		if err != nil {
			return nil, err
		}

		if !exist {
			break
		}
	}

	err := s.snapPaymentRepository.CreateSnapPayment(ctx, tx, snapPayment)
	if err != nil {
		return nil, err
	}

	invoiceURL, err := s.midtransService.CreateURLTransactionFromSaleInvoice(ctx, saleInvoice, snapPayment.ExternalID)
	if err != nil {
		return nil, err
	}

	snapPayment.InvoiceURL = null.StringFrom(invoiceURL)
	err = s.snapPaymentRepository.UpdateSnapPayment(ctx, tx, snapPayment)
	if err != nil {
		return nil, err
	}

	return snapPayment, nil
}
