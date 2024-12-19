package entity

import (
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"github.com/sherwin-77/go-tix/internal/enum"
	"gorm.io/gorm"
	"time"
)

type SnapPayment struct {
	BaseEntity
	UserID        uuid.UUID    `json:"user_id" gorm:"type:uuid;not null"`
	SaleInvoiceID uuid.UUID    `json:"sale_invoice_id" gorm:"type:uuid;not null"`
	ExternalID    string       `json:"external_id" gorm:"type:varchar(255);not null;uniqueIndex"`
	Amount        float64      `json:"amount" gorm:"type:decimal(16,2);not null"`
	SnapToken     null.String  `json:"snap_token" gorm:"type:varchar(255);"`
	InvoiceURL    null.String  `json:"invoice_url" gorm:"type:varchar(255);"`
	Status        string       `json:"status" gorm:"type:varchar(255);not null"`
	Method        string       `json:"method" gorm:"type:varchar(255);not null"`
	TransactionAt null.Time    `json:"transaction_at" gorm:"type:timestamp(6) with time zone"`
	ExpiredAt     null.Time    `json:"expired_at" gorm:"type:timestamp(6) with time zone"`
	CompletedAt   null.Time    `json:"completed_at" gorm:"type:timestamp(6) with time zone"`
	User          *User        `json:"user,omitempty" gorm:"foreignKey:UserID"`
	SaleInvoice   *SaleInvoice `json:"sale_invoice,omitempty" gorm:"foreignKey:SaleInvoiceID"`
}

func (s *SnapPayment) BeforeCreate(tx *gorm.DB) error {
	err := s.BaseEntity.BeforeCreate(tx)
	if err != nil {
		return err
	}

	if s.TransactionAt.IsZero() {
		s.TransactionAt = null.TimeFrom(time.Now())
	}

	return nil
}

func (s *SnapPayment) BeforeUpdate(tx *gorm.DB) error {
	switch s.Status {
	case string(enum.SnapPaymentStatusExpired):
		if s.ExpiredAt.IsZero() {
			s.ExpiredAt = null.TimeFrom(time.Now())
		}

	case string(enum.SnapPaymentStatusCompleted):
		if s.CompletedAt.IsZero() {
			s.CompletedAt = null.TimeFrom(time.Now())
		}

	case string(enum.SnapPaymentStatusCanceled):
		if s.CompletedAt.IsZero() {
			s.CompletedAt = null.TimeFrom(time.Now())
		}

	case string(enum.SnapPaymentStatusRejected):
		if s.CompletedAt.IsZero() {
			s.CompletedAt = null.TimeFrom(time.Now())
		}
	}

	return nil
}
