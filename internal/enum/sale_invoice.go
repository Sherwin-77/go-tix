package enum

type SaleInvoiceStatus string

const (
	SaleInvoiceStatusPending   SaleInvoiceStatus = "pending"
	SaleInvoiceStatusCompleted SaleInvoiceStatus = "completed"
	SaleInvoiceStatusCanceled  SaleInvoiceStatus = "canceled"
	SaleInvoiceStatusExpired   SaleInvoiceStatus = "expired"
	SaleInvoiceStatusRejected  SaleInvoiceStatus = "rejected"
	SaleInvoiceStatusRefunded  SaleInvoiceStatus = "refunded"
)
