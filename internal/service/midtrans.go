package service

import (
	"context"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/sherwin-77/go-tix/configs"
	"github.com/sherwin-77/go-tix/internal/entity"
	"github.com/sherwin-77/go-tix/pkg/payments"
	"math"
	"strings"
)

type MidtransService interface {
	CreateURLTransactionFromSaleInvoice(ctx context.Context, saleInvoice *entity.SaleInvoice, orderID string) (string, error)
}

type midtransService struct {
	midtransConfig configs.MidtransConfig
}

func NewMidtransService(midtransConfig configs.MidtransConfig) MidtransService {
	return &midtransService{
		midtransConfig: midtransConfig,
	}
}

func (s *midtransService) newSnapRequest(saleInvoice *entity.SaleInvoice, orderID string) *snap.Request {
	metadata := saleInvoice.Metadata.Data()
	parts := strings.SplitAfterN(metadata.FullName, " ", 2)
	var firstName, lastName string
	if len(parts) > 0 {
		firstName = parts[0]
	}
	if len(parts) > 1 {
		lastName = parts[1]
	}

	snapRequest := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(saleInvoice.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: firstName,
			LName: lastName,
			Email: metadata.Email,
			Phone: metadata.PhoneNumber,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items:           &[]midtrans.ItemDetails{},
	}

	for _, item := range saleInvoice.SaleInvoiceItems {
		itemMetadata := item.Metadata.Data()
		*snapRequest.Items = append(*snapRequest.Items, midtrans.ItemDetails{
			ID:    "Ticket-" + item.InvoiceableID.String(),
			Name:  itemMetadata.Name,
			Qty:   int32(item.Qty),
			Price: int64(math.Round(item.Price)),
		})
	}

	return snapRequest
}

func (s *midtransService) CreateURLTransactionFromSaleInvoice(ctx context.Context, saleInvoice *entity.SaleInvoice, orderID string) (string, error) {
	client := payments.NewMidtransSnapClient(s.midtransConfig)
	client.Options.SetContext(ctx)

	response, err := client.CreateTransactionUrl(s.newSnapRequest(saleInvoice, orderID))
	if err != nil {
		return "", err
	}

	return response, nil
}
