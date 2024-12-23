package entity

import (
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"github.com/sherwin-77/go-tix/internal/enum"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type SaleInvoiceMetadata struct {
	FullName           string      `json:"full_name"`
	PhoneNumber        string      `json:"phone_number"`
	Email              string      `json:"email"`
	IdentityCardNumber null.String `json:"identity_card_number"`
}

type SaleInvoice struct {
	BaseEntity
	UserID        uuid.UUID                               `json:"user_id" gorm:"type:uuid;not null"`
	Number        string                                  `json:"number" gorm:"type:varchar(50);uniqueIndex"`
	Subtotal      float64                                 `json:"subtotal" gorm:"type:decimal(16,2);not null"`
	ServiceFee    float64                                 `json:"service_fee" gorm:"type:decimal(16,2);not null"`
	PaymentFee    float64                                 `json:"payment_fee" gorm:"type:decimal(16,2);not null"`
	Discount      float64                                 `json:"discount" gorm:"type:decimal(16,2);not null"`
	Vat           float64                                 `json:"vat" gorm:"type:decimal(16,2);not null"`
	Total         float64                                 `json:"total" gorm:"type:decimal(16,2);not null"`
	Status        string                                  `json:"status" gorm:"type:varchar(50);not null"`
	TransactionAt null.Time                               `json:"transaction_at" gorm:"type:timestamp(6) with time zone"`
	DueAt         null.Time                               `json:"due_at" gorm:"type:timestamp(6) with time zone"`
	CompletedAt   null.Time                               `json:"completed_at" gorm:"type:timestamp(6) with time zone"`
	CanceledAt    null.Time                               `json:"canceled_at" gorm:"type:timestamp(6) with time zone"`
	ExpiredAt     null.Time                               `json:"expired_at" gorm:"type:timestamp(6) with time zone"`
	RejectedAt    null.Time                               `json:"rejected_at" gorm:"type:timestamp(6) with time zone"`
	RefundedAt    null.Time                               `json:"refunded_at" gorm:"type:timestamp(6) with time zone"`
	Metadata      datatypes.JSONType[SaleInvoiceMetadata] `json:"metadata" gorm:"type:jsonb"`

	User             *User              `json:"user,omitempty" gorm:"foreignKey:UserID"`
	SaleInvoiceItems []*SaleInvoiceItem `json:"sale_invoice_items,omitempty" gorm:"foreignKey:SaleInvoiceID"`
	SnapPayments     []*SnapPayment     `json:"snap_payments,omitempty" gorm:"foreignKey:SaleInvoiceID"`
}

func (s *SaleInvoice) BeforeCreate(tx *gorm.DB) error {
	err := s.BaseEntity.BeforeCreate(tx)
	if err != nil {
		return err
	}

	if s.TransactionAt.IsZero() {
		s.TransactionAt = null.TimeFrom(time.Now())
	}

	if s.DueAt.IsZero() {
		// Default to 24 hours if not set
		s.DueAt = null.TimeFrom(s.TransactionAt.ValueOrZero().Add(time.Hour * 24))
	}

	return nil
}

func (s *SaleInvoice) BeforeUpdate(tx *gorm.DB) error {
	switch s.Status {
	case string(enum.SaleInvoiceStatusCompleted):
		if s.CompletedAt.IsZero() {
			s.CompletedAt = null.TimeFrom(time.Now())
		}

	case string(enum.SaleInvoiceStatusCanceled):
		if s.CanceledAt.IsZero() {
			s.CanceledAt = null.TimeFrom(time.Now())
		}

	case string(enum.SaleInvoiceStatusExpired):
		if s.ExpiredAt.IsZero() {
			s.ExpiredAt = null.TimeFrom(time.Now())
		}

	case string(enum.SaleInvoiceStatusRejected):
		if s.RejectedAt.IsZero() {
			s.RejectedAt = null.TimeFrom(time.Now())
		}

	case string(enum.SaleInvoiceStatusRefunded):
		if s.RefundedAt.IsZero() {
			s.RefundedAt = null.TimeFrom(time.Now())
		}
	}

	return nil
}
