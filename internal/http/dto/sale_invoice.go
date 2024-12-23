package dto

import (
	"github.com/guregu/null/v5"
	"github.com/sherwin-77/go-tix/internal/domain"
	"github.com/sherwin-77/go-tix/internal/entity"
)

/* -------------------------------------------------------------------------- */
/*                                   Request                                  */
/* -------------------------------------------------------------------------- */

type CheckoutRequest struct {
	FullName           string                `json:"full_name" validate:"required"`
	PhoneNumber        string                `json:"phone_number" validate:"required,e164"`
	Email              string                `json:"email" validate:"required,email"`
	IdentityCardNumber *string               `json:"identity_card_number"`
	Items              []CheckoutRequestItem `json:"items" validate:"required,gte=1,dive"`
}

type CheckoutRequestItem struct {
	TicketID string `json:"ticket_id" validate:"required,uuid"`
	Qty      int    `json:"qty" validate:"required,gte=1,lte=5"` // Notice hardcode limit, change if needed
}

/* -------------------------------------------------------------------------- */
/*                                  Response                                  */
/* -------------------------------------------------------------------------- */

type SaleInvoiceListResponse struct {
	ID            string    `json:"id"`
	Number        string    `json:"number"`
	Subtotal      float64   `json:"subtotal"`
	Total         float64   `json:"total"`
	Status        string    `json:"status"`
	TransactionAt null.Time `json:"transaction_at"`
	DueAt         null.Time `json:"due_at"`
	CompletedAt   null.Time `json:"completed_at"`
	CanceledAt    null.Time `json:"canceled_at"`
	ExpiredAt     null.Time `json:"expired_at"`
	RejectedAt    null.Time `json:"rejected_at"`
	RefundedAt    null.Time `json:"refunded_at"`
}

type SaleInvoiceResponse struct {
	ID            string    `json:"id"`
	Number        string    `json:"number"`
	Subtotal      float64   `json:"subtotal"`
	ServiceFee    float64   `json:"service_fee"`
	PaymentFee    float64   `json:"payment_fee"`
	Discount      float64   `json:"discount"`
	Vat           float64   `json:"vat"`
	Total         float64   `json:"total"`
	Status        string    `json:"status"`
	TransactionAt null.Time `json:"transaction_at"`
	DueAt         null.Time `json:"due_at"`
	CompletedAt   null.Time `json:"completed_at"`
	CanceledAt    null.Time `json:"canceled_at"`
	ExpiredAt     null.Time `json:"expired_at"`
	RejectedAt    null.Time `json:"rejected_at"`
	RefundedAt    null.Time `json:"refunded_at"`
	UserInfo      struct {
		FullName           string      `json:"full_name"`
		PhoneNumber        string      `json:"phone_number"`
		Email              string      `json:"email"`
		IdentityCardNumber null.String `json:"identity_card_number"`
	} `json:"user_info"`
	SaleInvoiceItems []*SaleInvoiceItemResponse `json:"sale_invoice_items,omitempty"`
}

type BillResponse struct {
	Subtotal   float64            `json:"subtotal"`
	ServiceFee float64            `json:"service_fee"`
	PaymentFee float64            `json:"payment_fee"`
	Discount   float64            `json:"discount"`
	Vat        float64            `json:"vat"`
	Total      float64            `json:"total"`
	Items      []BillItemResponse `json:"items"`
}

type BillItemResponse struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
	Total float64 `json:"total"`
}

type CheckoutResponse struct {
	InvoiceURL string    `json:"invoice_url"`
	ExpiredAt  null.Time `json:"expired_at"`
}

func NewSaleInvoiceListResponse(saleInvoices []entity.SaleInvoice) []SaleInvoiceListResponse {
	response := make([]SaleInvoiceListResponse, 0)

	for _, saleInvoice := range saleInvoices {
		response = append(response, SaleInvoiceListResponse{
			ID:            saleInvoice.ID.String(),
			Number:        saleInvoice.Number,
			Subtotal:      saleInvoice.Subtotal,
			Total:         saleInvoice.Total,
			Status:        saleInvoice.Status,
			TransactionAt: saleInvoice.TransactionAt,
			DueAt:         saleInvoice.DueAt,
			CompletedAt:   saleInvoice.CompletedAt,
			CanceledAt:    saleInvoice.CanceledAt,
			ExpiredAt:     saleInvoice.ExpiredAt,
			RejectedAt:    saleInvoice.RejectedAt,
			RefundedAt:    saleInvoice.RefundedAt,
		})
	}

	return response
}

func NewSaleInvoiceResponse(saleInvoice *entity.SaleInvoice) *SaleInvoiceResponse {
	metadata := saleInvoice.Metadata.Data()
	response := &SaleInvoiceResponse{
		ID:            saleInvoice.ID.String(),
		Number:        saleInvoice.Number,
		Subtotal:      saleInvoice.Subtotal,
		ServiceFee:    saleInvoice.ServiceFee,
		PaymentFee:    saleInvoice.PaymentFee,
		Discount:      saleInvoice.Discount,
		Vat:           saleInvoice.Vat,
		Total:         saleInvoice.Total,
		Status:        saleInvoice.Status,
		TransactionAt: saleInvoice.TransactionAt,
		DueAt:         saleInvoice.DueAt,
		CompletedAt:   saleInvoice.CompletedAt,
		CanceledAt:    saleInvoice.CanceledAt,
		ExpiredAt:     saleInvoice.ExpiredAt,
		RejectedAt:    saleInvoice.RejectedAt,
		RefundedAt:    saleInvoice.RefundedAt,
		UserInfo: struct {
			FullName           string      `json:"full_name"`
			PhoneNumber        string      `json:"phone_number"`
			Email              string      `json:"email"`
			IdentityCardNumber null.String `json:"identity_card_number"`
		}{
			FullName:           metadata.FullName,
			PhoneNumber:        metadata.PhoneNumber,
			Email:              metadata.Email,
			IdentityCardNumber: metadata.IdentityCardNumber,
		},
	}

	for _, item := range saleInvoice.SaleInvoiceItems {
		response.SaleInvoiceItems = append(response.SaleInvoiceItems, NewSaleInvoiceItemResponse(item))
	}

	return response
}

func NewBillResponseFromInvoicePricing(pricing *domain.InvoicePricing) *BillResponse {
	response := &BillResponse{
		Subtotal:   pricing.Subtotal,
		ServiceFee: pricing.ServiceFee,
		PaymentFee: pricing.PaymentFee,
		Discount:   pricing.Discount,
		Vat:        pricing.Vat,
		Total:      pricing.Total,
	}

	for _, item := range pricing.InvoiceItems {
		response.Items = append(response.Items, BillItemResponse{
			Name:  item.Name,
			Price: item.Price,
			Qty:   item.Qty,
			Total: item.Total,
		})
	}

	return response
}

func NewCheckoutResponseFromCheckoutData(checkoutData *domain.CheckoutData) *CheckoutResponse {
	return &CheckoutResponse{
		InvoiceURL: checkoutData.InvoiceURL,
		ExpiredAt:  checkoutData.ExpiredAt,
	}
}
