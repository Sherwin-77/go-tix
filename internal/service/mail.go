package service

import (
	"fmt"
	"github.com/sherwin-77/go-tix/configs"
	"github.com/sherwin-77/go-tix/internal/domain"
	"github.com/sherwin-77/go-tix/internal/entity"
	"github.com/wneessen/go-mail"
	"html/template"
	"strconv"
)

type MailService struct {
	mailConfig configs.MailConfig
	client     *mail.Client
}

func NewMailService(mailConfig configs.MailConfig) MailService {
	port, err := strconv.Atoi(mailConfig.Port)
	if err != nil {
		panic(err)
	}

	client, err := mail.NewClient(
		mailConfig.Host,
		mail.WithSMTPAuth(mail.SMTPAuthLogin),
		mail.WithUsername(mailConfig.Username),
		mail.WithPassword(mailConfig.Password),
		mail.WithPort(port),
		mail.WithSSLPort(false),
	)
	if err != nil {
		panic(err)
	}

	return MailService{
		mailConfig: mailConfig,
		client:     client,
	}
}

func (m *MailService) SendTicket(email string, saleInvoice *entity.SaleInvoice) error {
	message := mail.NewMsg()
	if err := message.From(m.mailConfig.FromAddress); err != nil {
		return err
	}
	if err := message.AddTo(email); err != nil {
		return err
	}
	message.Subject(fmt.Sprintf("Go Tix - Invoice #%s", saleInvoice.Number))

	var ticketDatas []domain.TicketData
	for _, invoiceItem := range saleInvoice.SaleInvoiceItems {
		itemMetadata := invoiceItem.Metadata.Data()
		for _, code := range itemMetadata.Codes {
			ticketDatas = append(ticketDatas, domain.TicketData{
				Name: itemMetadata.Name,
				Code: code,
			})
		}
	}

	metadata := saleInvoice.Metadata.Data()
	mailData := map[string]interface{}{
		"Number":   saleInvoice.Number,
		"Date":     saleInvoice.CompletedAt.ValueOrZero(),
		"Customer": metadata.FullName,
		"Total":    saleInvoice.Total,
		"Items":    ticketDatas,
	}

	tmpl, err := template.ParseFiles("views/mail-ticket.html")
	if err != nil {
		return err
	}

	if err = message.SetBodyHTMLTemplate(tmpl, mailData); err != nil {
		return err
	}

	return m.client.DialAndSend(message)
}
