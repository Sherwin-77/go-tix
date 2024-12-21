package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-tix/internal/domain"
	"github.com/sherwin-77/go-tix/internal/entity"
	"github.com/sherwin-77/go-tix/internal/enum"
	"github.com/sherwin-77/go-tix/internal/http/dto"
	"github.com/sherwin-77/go-tix/internal/repository"
	"github.com/sherwin-77/go-tix/pkg/query"
	"github.com/sherwin-77/go-tix/pkg/response"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type SaleInvoiceService interface {
	GetSaleInvoices(ctx context.Context, queryParams url.Values) ([]entity.SaleInvoice, *response.Meta, error)
	GetSaleInvoiceByID(ctx context.Context, id string) (*entity.SaleInvoice, error)

	GetUserSaleInvoices(ctx context.Context, queryParams url.Values, userID string) ([]entity.SaleInvoice, *response.Meta, error)
	GetUserSaleInvoiceByID(ctx context.Context, id string, userID string) (*entity.SaleInvoice, error)
	Bill(ctx context.Context, request dto.CheckoutRequest) (*domain.InvoicePricing, error)
	Checkout(ctx context.Context, request dto.CheckoutRequest, userID string) (*domain.CheckoutData, error)
}

type saleInvoiceService struct {
	saleInvoiceRepository repository.SaleInvoiceRepository
	ticketRepository      repository.TicketRepository
	transactionService    TransactionService
	saleInvoiceBuilder    query.Builder
	randomizer            *rand.Rand
}

func NewSaleInvoiceService(
	saleInvoiceRepository repository.SaleInvoiceRepository,
	ticketRepository repository.TicketRepository,
	transactionService TransactionService,
	saleInvoiceBuilder query.Builder,
) SaleInvoiceService {
	return &saleInvoiceService{
		saleInvoiceRepository,
		ticketRepository,
		transactionService,
		saleInvoiceBuilder,
		rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *saleInvoiceService) GetSaleInvoices(ctx context.Context, queryParams url.Values) ([]entity.SaleInvoice, *response.Meta, error) {
	var saleInvoices []entity.SaleInvoice
	var err error
	var meta *response.Meta
	db := s.saleInvoiceRepository.SingleTransaction()

	db, meta = s.saleInvoiceBuilder.ApplyBuilder(db, queryParams, &entity.SaleInvoice{})

	saleInvoices, err = s.saleInvoiceRepository.GetSaleInvoices(ctx, db)
	if err != nil {
		return nil, nil, err
	}

	return saleInvoices, meta, nil
}

func (s *saleInvoiceService) GetSaleInvoiceByID(ctx context.Context, id string) (*entity.SaleInvoice, error) {
	db := s.saleInvoiceRepository.SingleTransaction()
	db = s.saleInvoiceRepository.WithPreloads(db, map[string][]interface{}{"User": nil, "SaleInvoiceItems.Ticket": nil})

	saleInvoice, err := s.saleInvoiceRepository.GetSaleInvoiceByID(ctx, db, id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Invoice not found")
	}

	return saleInvoice, nil
}

func (s *saleInvoiceService) GetUserSaleInvoices(ctx context.Context, queryParams url.Values, userID string) ([]entity.SaleInvoice, *response.Meta, error) {
	var saleInvoices []entity.SaleInvoice
	var err error
	var meta *response.Meta
	db := s.saleInvoiceRepository.SingleTransaction()

	db, meta = s.saleInvoiceBuilder.ApplyBuilder(db, queryParams, &entity.SaleInvoice{})

	saleInvoices, err = s.saleInvoiceRepository.GetUserSaleInvoices(ctx, db, userID)
	if err != nil {
		return nil, nil, err
	}

	return saleInvoices, meta, nil
}

func (s *saleInvoiceService) GetUserSaleInvoiceByID(ctx context.Context, id string, userID string) (*entity.SaleInvoice, error) {
	saleInvoice, err := s.GetSaleInvoiceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if saleInvoice.UserID.String() != userID {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Invoice not found")
	}

	return saleInvoice, nil
}

func (s *saleInvoiceService) Bill(ctx context.Context, request dto.CheckoutRequest) (*domain.InvoicePricing, error) {
	invoicePricing := &domain.InvoicePricing{
		ServiceFee: 0,
		PaymentFee: 0,
		Discount:   0,
		Vat:        0,
	}

	ticketAmounts := make(map[string]int)
	var ticketIDs []string

	for _, item := range request.Items {
		ticketAmounts[item.TicketID] += item.Qty
		ticketIDs = append(ticketIDs, item.TicketID)
	}

	db := s.saleInvoiceRepository.SingleTransaction()
	db = s.saleInvoiceRepository.WithPreloads(db, map[string][]interface{}{"Event": nil})
	tickets, err := s.ticketRepository.GetTicketsByTicketIDs(ctx, db, ticketIDs)
	if err != nil {
		return nil, err
	}

	if len(tickets) != len(ticketAmounts) {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Ticket not found")
	}

	for _, ticket := range tickets {
		invoicePricing.Subtotal += ticket.Price * float64(ticketAmounts[ticket.ID.String()])
		invoicePricing.InvoiceItems = append(invoicePricing.InvoiceItems, domain.InvoiceItemPricing{
			ID:    ticket.ID,
			Name:  ticket.Category,
			Price: ticket.Price,
			Qty:   ticketAmounts[ticket.ID.String()],
			Total: ticket.Price * float64(ticketAmounts[ticket.ID.String()]),
		})

		endAt := time.Time(ticket.Event.EndAt)
		if endAt.Before(time.Now()) {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "Event has ended")
		}
	}
	invoicePricing.Total = invoicePricing.Subtotal + invoicePricing.ServiceFee + invoicePricing.PaymentFee - invoicePricing.Discount + invoicePricing.Vat

	return invoicePricing, nil
}

func (s *saleInvoiceService) Checkout(ctx context.Context, request dto.CheckoutRequest, userID string) (*domain.CheckoutData, error) {
	data := &domain.CheckoutData{}

	err := s.saleInvoiceRepository.WithTransaction(func(tx *gorm.DB) error {
		invoicePricing, err := s.Bill(ctx, request)
		if err != nil {
			return err
		}

		saleInvoice := &entity.SaleInvoice{
			Status:     string(enum.SaleInvoiceStatusPending),
			UserID:     uuid.MustParse(userID),
			Subtotal:   invoicePricing.Subtotal,
			ServiceFee: invoicePricing.ServiceFee,
			PaymentFee: invoicePricing.PaymentFee,
			Discount:   invoicePricing.Discount,
			Vat:        invoicePricing.Vat,
			Total:      invoicePricing.Total,
			Metadata: datatypes.NewJSONType(entity.SaleInvoiceMetadata{
				FullName:           request.FullName,
				PhoneNumber:        request.PhoneNumber,
				Email:              request.Email,
				IdentityCardNumber: null.StringFromPtr(request.IdentityCardNumber),
			}),
		}

		for _, item := range invoicePricing.InvoiceItems {
			invoiceItem := &entity.SaleInvoiceItem{
				InvoiceableID:   item.ID,
				InvoiceableType: "tickets",
				Price:           item.Price,
				Qty:             item.Qty,
				Total:           item.Total,
			}
			metadata := entity.SaleInvoiceItemMetadata{
				Name:  item.Name,
				Codes: []string{},
			}

			for i := 0; i < item.Qty; i++ {
				metadata.Codes = append(metadata.Codes, uuid.New().String())
			}

			invoiceItem.Metadata = datatypes.NewJSONType(metadata)
			saleInvoice.SaleInvoiceItems = append(saleInvoice.SaleInvoiceItems, invoiceItem)
		}

		err = s.saleInvoiceRepository.CreateSaleInvoice(ctx, tx, saleInvoice)
		if err != nil {
			return err
		}

		snapPayment, err := s.transactionService.CreateTransaction(ctx, tx, saleInvoice)
		if err != nil {
			return err
		}

		saleInvoice.Number = snapPayment.ExternalID
		if err = s.saleInvoiceRepository.UpdateSaleInvoice(ctx, tx, saleInvoice); err != nil {
			return err
		}

		data = &domain.CheckoutData{
			InvoiceURL: snapPayment.InvoiceURL.ValueOrZero(),
			ExpiredAt:  saleInvoice.DueAt,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}
