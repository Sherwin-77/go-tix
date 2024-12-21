package domain

import (
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
)

type InvoicePricing struct {
	Subtotal     float64
	ServiceFee   float64
	PaymentFee   float64
	Discount     float64
	Vat          float64
	Total        float64
	InvoiceItems []InvoiceItemPricing
}

type InvoiceItemPricing struct {
	ID    uuid.UUID
	Name  string
	Price float64
	Qty   int
	Total float64
}

type CheckoutData struct {
	InvoiceURL string
	ExpiredAt  null.Time
}
