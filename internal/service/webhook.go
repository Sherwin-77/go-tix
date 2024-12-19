package service

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-tix/configs"
	"github.com/sherwin-77/go-tix/internal/enum"
	"github.com/sherwin-77/go-tix/internal/repository"
	"github.com/sherwin-77/go-tix/pkg/payments"
	"gorm.io/gorm"
	"net/http"
)

type WebhookService interface {
	HandleMidtransNotification(ctx context.Context, payload map[string]interface{}) error
}

type webhookService struct {
	midtransConfig        configs.MidtransConfig
	saleInvoiceRepository repository.SaleInvoiceRepository
	snapPaymentRepository repository.SnapPaymentRepository
}

func NewWebhookService(
	midtransConfig configs.MidtransConfig,
	saleInvoiceRepository repository.SaleInvoiceRepository,
	snapPaymentRepository repository.SnapPaymentRepository,
) WebhookService {
	return &webhookService{
		midtransConfig,
		saleInvoiceRepository,
		snapPaymentRepository,
	}
}

func (s *webhookService) HandleMidtransNotification(ctx context.Context, payload map[string]interface{}) error {
	client := payments.NewMidTransCoreApiClient(s.midtransConfig)
	client.Options.SetContext(ctx)

	return s.snapPaymentRepository.WithTransaction(func(tx *gorm.DB) error {
		orderID, exists := payload["order_id"].(string)
		if !exists {
			return echo.NewHTTPError(http.StatusBadRequest, "Order ID does not exist")
		}

		snapPayment, err := s.snapPaymentRepository.GetByExternalID(ctx, tx, orderID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		saleInvoice, err := s.saleInvoiceRepository.GetSaleInvoiceByID(ctx, tx, snapPayment.SaleInvoiceID.String())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		transactionStatusResp, e := client.CheckTransaction(orderID)
		if e != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if transactionStatusResp.TransactionStatus == "capture" {
			if transactionStatusResp.FraudStatus == "challenge" {
				// Challenge status
			} else if transactionStatusResp.FraudStatus == "accept" {
				snapPayment.Status = string(enum.SnapPaymentStatusCompleted)
				saleInvoice.Status = string(enum.SaleInvoiceStatusCompleted)
			}
		} else if transactionStatusResp.TransactionStatus == "settlement" {
			snapPayment.Status = string(enum.SnapPaymentStatusCompleted)
			saleInvoice.Status = string(enum.SaleInvoiceStatusCompleted)
		} else if transactionStatusResp.TransactionStatus == "deny" {
			// Ignore for the time being. Because most of the time it allows payment retries
		} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
			snapPayment.Status = string(enum.SnapPaymentStatusRejected)

			if transactionStatusResp.TransactionStatus == "cancel" {
				saleInvoice.Status = string(enum.SaleInvoiceStatusCanceled)
			} else {
				saleInvoice.Status = string(enum.SaleInvoiceStatusExpired)
			}
		} else if transactionStatusResp.TransactionStatus == "pending" {
			snapPayment.Status = string(enum.SnapPaymentStatusPending)
			saleInvoice.Status = string(enum.SaleInvoiceStatusPending)
		}

		err = s.snapPaymentRepository.UpdateSnapPayment(ctx, tx, snapPayment)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		err = s.saleInvoiceRepository.UpdateSaleInvoice(ctx, tx, saleInvoice)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return nil
	})
}
