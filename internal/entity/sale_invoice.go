package entity

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"time"
)

type SaleInvoiceMetadata struct {
}

type SaleInvoice struct {
	BaseEntity
	UserID        uuid.UUID                               `json:"user_id" gorm:"type:uuid;not null"`
	Number        string                                  `json:"number" gorm:"type:varchar(50);"`
	Subtotal      float64                                 `json:"subtotal" gorm:"type:decimal(16,2);not null"`
	ServiceFee    float64                                 `json:"service_fee" gorm:"type:decimal(16,2);not null"`
	PaymentFee    float64                                 `json:"payment_fee" gorm:"type:decimal(16,2);not null"`
	Discount      float64                                 `json:"discount" gorm:"type:decimal(16,2);not null"`
	Vat           float64                                 `json:"vat" gorm:"type:decimal(16,2);not null"`
	Total         float64                                 `json:"total" gorm:"type:decimal(16,2);not null"`
	Status        string                                  `json:"status" gorm:"type:varchar(50);not null"`
	TransactionAt time.Time                               `json:"transaction_at" gorm:"type:timestamp(6) with time zone"`
	DueAt         time.Time                               `json:"due_at" gorm:"type:timestamp(6) with time zone"`
	CompletedAt   time.Time                               `json:"completed_at" gorm:"type:timestamp(6) with time zone"`
	CanceledAt    time.Time                               `json:"canceled_at" gorm:"type:timestamp(6) with time zone"`
	ExpiredAt     time.Time                               `json:"expired_at" gorm:"type:timestamp(6) with time zone"`
	RejectedAt    time.Time                               `json:"rejected_at" gorm:"type:timestamp(6) with time zone"`
	RefundedAt    time.Time                               `json:"refunded_at" gorm:"type:timestamp(6) with time zone"`
	Metadata      datatypes.JSONType[SaleInvoiceMetadata] `json:"metadata" gorm:"type:jsonb"`

	User             *User              `json:"user,omitempty" gorm:"foreignKey:UserID"`
	SaleInvoiceItems []*SaleInvoiceItem `json:"sale_invoice_items,omitempty" gorm:"foreignKey:SaleInvoiceID"`
}
