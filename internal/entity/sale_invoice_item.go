package entity

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type SaleInvoiceItemMetadata struct {
	Name string `json:"name"`
}

type SaleInvoiceItem struct {
	BaseEntity
	SaleInvoiceID   uuid.UUID                                   `json:"sale_invoice_id" gorm:"type:uuid;not null"`
	InvoiceableID   uuid.UUID                                   `json:"invoiceable_id" gorm:"type:uuid;not null"`
	InvoiceableType string                                      `json:"invoiceable_type" gorm:"type:varchar(255);not null"`
	Price           float64                                     `json:"price" gorm:"type:decimal(12,2);not null"`
	Qty             int                                         `json:"qty" gorm:"type:integer;not null"`
	Total           float64                                     `json:"total" gorm:"type:decimal(16,2);not null"`
	Metadata        datatypes.JSONType[SaleInvoiceItemMetadata] `json:"metadata" gorm:"type:jsonb"`

	Ticket *Ticket `json:"ticket,omitempty" gorm:"foreignKey:InvoiceableID"`
	// TODO invoiceable foreign
}
