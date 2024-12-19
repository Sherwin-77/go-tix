package entity

import "github.com/google/uuid"

type Ticket struct {
	BaseEntity
	EventID  uuid.UUID `json:"event_id" gorm:"type:uuid;not null"`
	Category string    `json:"category" gorm:"type:varchar(100);not null"`
	Price    float64   `json:"price" gorm:"type:decimal(12,2);not null"`

	Event            *Event             `json:"event,omitempty" gorm:"foreignKey:EventID"`
	SaleInvoiceItems []*SaleInvoiceItem `json:"sale_invoice_items,omitempty" gorm:"polymorphicId:InvoiceableID;polymorphicType:InvoiceableType;polymorphicValue:tickets"`
}
