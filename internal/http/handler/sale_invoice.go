package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-tix/internal/http/dto"
	"github.com/sherwin-77/go-tix/internal/service"
	"github.com/sherwin-77/go-tix/pkg/response"
	"net/http"
)

type SaleInvoiceHandler struct {
	saleInvoiceService service.SaleInvoiceService
}

func NewSaleInvoiceHandler(saleInvoiceService service.SaleInvoiceService) SaleInvoiceHandler {
	return SaleInvoiceHandler{
		saleInvoiceService,
	}
}

// Bill
//
//	@Summary	Bill
//	@Tags		[User] Sale Invoice
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.CheckoutRequest	true	"Checkout Request"
//	@Success	200		{object}	response.Response{data=dto.BillResponse}
//	@Router		/bill [post]
func (h *SaleInvoiceHandler) Bill(ctx echo.Context) error {
	var request dto.CheckoutRequest
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	if err := ctx.Validate(request); err != nil {
		return err
	}

	invoicePricing, err := h.saleInvoiceService.Bill(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	billResponse := dto.NewBillResponseFromInvoicePricing(invoicePricing)
	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", billResponse, nil))
}

// Checkout
//
//	@Summary	Checkout
//	@Tags		[User] Sale Invoice
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.CheckoutRequest	true	"Checkout Request"
//	@Success	200		{object}	response.Response{data=dto.CheckoutResponse}
//	@Router		/checkout [post]
func (h *SaleInvoiceHandler) Checkout(ctx echo.Context) error {
	userID := ctx.Get("user_id").(string)
	var request dto.CheckoutRequest
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	if err := ctx.Validate(request); err != nil {
		return err
	}

	checkoutData, err := h.saleInvoiceService.Checkout(ctx.Request().Context(), request, userID)
	if err != nil {
		return err
	}

	checkoutResponse := dto.NewCheckoutResponseFromCheckoutData(checkoutData)

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", checkoutResponse, nil))
}

// GetUserSaleInvoices
//
//	@Summary	Get User Sale Invoices
//	@Tags		[User] Sale Invoice
//	@Produce	json
//	@Success	200	{object}	response.Response{data=[]dto.SaleInvoiceListResponse}
//	@Router		/sale-invoices [get]
func (h *SaleInvoiceHandler) GetUserSaleInvoices(ctx echo.Context) error {
	userID := ctx.Get("user_id").(string)
	queryParams := ctx.QueryParams()

	saleInvoices, meta, err := h.saleInvoiceService.GetUserSaleInvoices(ctx.Request().Context(), queryParams, userID)
	if err != nil {
		return err
	}

	saleInvoiceResponses := dto.NewSaleInvoiceListResponse(saleInvoices)

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", saleInvoiceResponses, meta))
}

// GetUserSaleInvoiceByID
//
//	@Summary	Get User Sale Invoice By ID
//	@Tags		[User] Sale Invoice
//	@Produce	json
//	@Param		id	path		string	true	"Sale Invoice ID"
//	@Success	200	{object}	response.Response{data=dto.SaleInvoiceResponse}
//	@Router		/sale-invoices/{id} [get]
func (h *SaleInvoiceHandler) GetUserSaleInvoiceByID(ctx echo.Context) error {
	userID := ctx.Get("user_id").(string)
	eventID := ctx.Param("id")

	if eventID == "" {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	saleInvoice, err := h.saleInvoiceService.GetUserSaleInvoiceByID(ctx.Request().Context(), eventID, userID)
	if err != nil {
		return err
	}

	saleInvoiceResponse := dto.NewSaleInvoiceResponse(saleInvoice)

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", saleInvoiceResponse, nil))
}
