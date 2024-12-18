package service

import (
	"context"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-tix/internal/entity"
	"github.com/sherwin-77/go-tix/internal/enum"
	"github.com/sherwin-77/go-tix/internal/http/dto"
	"github.com/sherwin-77/go-tix/internal/repository"
	"github.com/sherwin-77/go-tix/pkg/query"
	"github.com/sherwin-77/go-tix/pkg/response"
	"gorm.io/gorm"
)

type EventApprovalService interface {
	GetEventApprovals(ctx context.Context, queryParams url.Values) ([]entity.EventApproval, *response.Meta, error)
	GetEventApprovalByID(ctx context.Context, id string) (*entity.EventApproval, error)
	HandleEventApproval(ctx context.Context, request dto.HandleEventApprovalRequest) error
	GetUserEventApprovals(ctx context.Context, queryParams url.Values) ([]entity.EventApproval, *response.Meta, error)
	GetEventApprovalForUserByID(ctx context.Context, id string) (*entity.EventApproval, error)
}

type eventApprovalService struct {
	eventService            EventService
	eventApprovalRepository repository.EventApprovalRepository
	eventApprovalBuilder    query.Builder
}

func NewEventApprovalService(
	eventService EventService,
	eventApprovalRepository repository.EventApprovalRepository,
	eventApprovalBuilder query.Builder,
) EventApprovalService {
	return &eventApprovalService{
		eventService,
		eventApprovalRepository,
		eventApprovalBuilder,
	}
}

/* -------------------------------------------------------------------------- */
/*                                Admin Service                               */
/* -------------------------------------------------------------------------- */

func (s *eventApprovalService) GetEventApprovals(ctx context.Context, queryParams url.Values) ([]entity.EventApproval, *response.Meta, error) {
	var eventApprovals []entity.EventApproval
	var err error
	var meta *response.Meta
	db := s.eventApprovalRepository.SingleTransaction()

	if queryParams.Get("status") != "" {
		db = db.Where("status = ?", queryParams.Get("status"))
	} else {
		db = db.Where("status = ?", enum.EventApprovalStatusPending)
	}

	db, meta = s.eventApprovalBuilder.ApplyBuilder(db, queryParams, &entity.EventApproval{})
	eventApprovals, err = s.eventApprovalRepository.GetEventApprovals(ctx, db)
	if err != nil {
		return nil, nil, err
	}

	return eventApprovals, meta, nil
}

func (s *eventApprovalService) GetEventApprovalByID(ctx context.Context, id string) (*entity.EventApproval, error) {
	db := s.eventApprovalRepository.SingleTransaction()
	db = s.eventApprovalRepository.WithPreloads(db, map[string][]interface{}{"EventApprovalTickets": nil})

	eventApproval, err := s.eventApprovalRepository.GetEventApprovalByID(ctx, db, id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Event approval not found")
	}

	return eventApproval, nil
}

func (s *eventApprovalService) HandleEventApproval(ctx context.Context, request dto.HandleEventApprovalRequest) error {
	return s.eventApprovalRepository.WithTransaction(func(tx *gorm.DB) error {
		db := s.eventApprovalRepository.WithPreloads(tx, map[string][]interface{}{"EventApprovalTickets": nil})

		eventApproval, err := s.eventApprovalRepository.GetEventApprovalByID(ctx, db, request.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Event approval not found")
		}

		if eventApproval.Status != string(enum.EventApprovalStatusPending) {
			return echo.NewHTTPError(http.StatusBadRequest, "Event approval is not pending")
		}

		if request.Action == "approve" {
			eventApproval.Status = string(enum.EventApprovalStatusApproved)

			if err = s.eventService.CreateEventFromEventApproval(ctx, eventApproval); err != nil {
				return err
			}
		} else if request.Action == "reject" {
			eventApproval.Status = string(enum.EventApprovalStatusRejected)
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid action")
		}

		if err = s.eventApprovalRepository.UpdateEventApproval(ctx, tx, eventApproval); err != nil {
			return err
		}

		return nil
	})
}

/* -------------------------------------------------------------------------- */
/*                                User Service                                */
/* -------------------------------------------------------------------------- */

func (s *eventApprovalService) GetUserEventApprovals(ctx context.Context, queryParams url.Values) ([]entity.EventApproval, *response.Meta, error) {
	var eventApprovals []entity.EventApproval
	var err error
	var meta *response.Meta
	db := s.eventApprovalRepository.SingleTransaction()

	if queryParams.Get("status") != "" {
		db = db.Where("status = ?", queryParams.Get("status"))
	} else {
		db = db.Where("status = ?", enum.EventApprovalStatusPending)
	}

	db, meta = s.eventApprovalBuilder.ApplyBuilder(db, queryParams, &entity.EventApproval{})
	eventApprovals, err = s.eventApprovalRepository.GetEventApprovals(ctx, db)
	if err != nil {
		return nil, nil, err
	}

	return eventApprovals, meta, nil
}

func (s *eventApprovalService) GetEventApprovalForUserByID(ctx context.Context, id string) (*entity.EventApproval, error) {
	db := s.eventApprovalRepository.SingleTransaction()
	db = s.eventApprovalRepository.WithPreloads(db, map[string][]interface{}{"EventApprovalTickets": nil})

	eventApproval, err := s.eventApprovalRepository.GetEventApprovalByID(ctx, db, id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Event approval not found")
	}

	return eventApproval, nil
}

// HandleEventApprovalForUser creates an event approval for a user with the given validation on admin.
func (s *eventApprovalService) HandleEventApprovalForUser(ctx context.Context, request dto.HandleEventApprovalRequest) error {
	return s.eventApprovalRepository.WithTransaction(func(tx *gorm.DB) error {
		db := s.eventApprovalRepository.WithPreloads(tx, map[string][]interface{}{"EventApprovalTickets": nil})

		eventApproval, err := s.eventApprovalRepository.GetEventApprovalByID(ctx, db, request.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Event approval not found")
		}

		if eventApproval.Status != string(enum.EventApprovalStatusPending) {
			return echo.NewHTTPError(http.StatusBadRequest, "Event approval is not pending")
		}

		if request.Action == "approve" {
			eventApproval.Status = string(enum.EventApprovalStatusApproved)

			if err = s.eventService.CreateEventFromEventApproval(ctx, eventApproval); err != nil {
				return err
			}
		} else if request.Action == "reject" {
			eventApproval.Status = string(enum.EventApprovalStatusRejected)
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid action")
		}

		if err = s.eventApprovalRepository.UpdateEventApproval(ctx, tx, eventApproval); err != nil {
			return err
		}

		return nil
	})
}
