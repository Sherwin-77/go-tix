package repository

import (
	"context"
	"github.com/sherwin-77/go-tix/internal/entity"
	"gorm.io/gorm"
)

type TicketRepository interface {
	GetTicketsByTicketIDs(ctx context.Context, tx *gorm.DB, ticketIDs []string) ([]entity.Ticket, error)
}

type ticketRepository struct {
	baseRepository
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{
		baseRepository: baseRepository{db},
	}
}

func (r ticketRepository) GetTicketsByTicketIDs(ctx context.Context, tx *gorm.DB, ticketIDs []string) ([]entity.Ticket, error) {
	var tickets []entity.Ticket

	if err := tx.WithContext(ctx).Where("id in (?)", ticketIDs).Find(&tickets).Error; err != nil {
		return nil, err
	}

	return tickets, nil
}
