package dto

import "github.com/sherwin-77/go-tix/internal/entity"

type SaleInvoiceItemResponse struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
	Total float64 `json:"total"`
}

func NewSaleInvoiceItemResponse(saleInvoiceItem *entity.SaleInvoiceItem) *SaleInvoiceItemResponse {
	metadata := saleInvoiceItem.Metadata.Data()
	return &SaleInvoiceItemResponse{
		Name:  metadata.Name,
		Price: saleInvoiceItem.Price,
		Qty:   saleInvoiceItem.Qty,
		Total: saleInvoiceItem.Total,
	}
}
