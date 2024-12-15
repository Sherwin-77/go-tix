package service

import (
	"context"
	"encoding/json"
	"github.com/guregu/null/v5"
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-tix/internal/entity"
	"github.com/sherwin-77/go-tix/internal/enum"
	"github.com/sherwin-77/go-tix/internal/http/dto"
	"github.com/sherwin-77/go-tix/internal/repository"
	"github.com/sherwin-77/go-tix/pkg/caches"
	"github.com/sherwin-77/go-tix/pkg/query"
	"github.com/sherwin-77/go-tix/pkg/response"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"time"
)

type EventService interface {
	GetEvents(ctx context.Context, queryParams url.Values) ([]entity.EventWithMinMaxPrice, *response.Meta, error)
	GetEventByID(ctx context.Context, id string) (*entity.Event, error)
	CreateEvent(ctx context.Context, request dto.CreateEventRequest) error
	UpdateEvent(ctx context.Context, request dto.UpdateEventRequest) error
	GetUserEvents(ctx context.Context, queryParams url.Values) ([]entity.EventWithMinMaxPrice, *response.Meta, error)
	GetUserEventByID(ctx context.Context, id string) (*entity.Event, error)
	RegisterEvent(ctx context.Context, request dto.CreateEventRequest, userID string) error
}

type eventService struct {
	userRepository          repository.UserRepository
	eventRepository         repository.EventRepository
	eventApprovalRepository repository.EventApprovalRepository
	eventBuilder            query.Builder
	cache                   caches.Cache
}

func NewEventService(
	userRepository repository.UserRepository,
	eventRepository repository.EventRepository,
	eventApprovalRepository repository.EventApprovalRepository,
	eventBuilder query.Builder,
	cache caches.Cache,
) EventService {
	return &eventService{
		userRepository,
		eventRepository,
		eventApprovalRepository,
		eventBuilder,
		cache,
	}
}

/* -------------------------------------------------------------------------- */
/*                                Admin Service                               */
/* -------------------------------------------------------------------------- */

func (s *eventService) GetEvents(ctx context.Context, queryParams url.Values) ([]entity.EventWithMinMaxPrice, *response.Meta, error) {
	var events []entity.EventWithMinMaxPrice
	var err error
	var meta *response.Meta
	db := s.eventRepository.SingleTransaction()

	db, meta = s.eventBuilder.ApplyBuilder(db, queryParams, &entity.EventWithMinMaxPrice{})
	if queryParams.Get("status") != "" {
		db = db.Where("status = ?", queryParams.Get("status"))
	}

	events, err = s.eventRepository.GetEventsWithMinMaxPrice(ctx, db)
	if err != nil {
		return nil, nil, err
	}

	return events, meta, nil
}

func (s *eventService) GetEventByID(ctx context.Context, id string) (*entity.Event, error) {
	eventKey := "event:" + id
	event := &entity.Event{}
	cachedData := s.cache.Get(eventKey)
	if cachedData != "" {
		if err := json.Unmarshal([]byte(cachedData), event); err != nil {
			return nil, err
		}
	} else {
		var err error
		db := s.eventRepository.SingleTransaction()
		db = s.eventRepository.WithPreloads(db, map[string][]interface{}{"Tickets": nil})

		event, err = s.eventRepository.GetEventByID(ctx, db, id)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusNotFound, "Event not found")
		}

		data, _ := json.Marshal(event)

		if err = s.cache.Set(eventKey, string(data), 5*time.Minute); err != nil {
			return nil, err
		}
	}

	return event, nil
}

func (s *eventService) CreateEvent(ctx context.Context, request dto.CreateEventRequest) error {
	return s.eventRepository.WithTransaction(func(tx *gorm.DB) error {
		event := entity.Event{
			Title:       request.Title,
			Description: null.StringFromPtr(request.Description),
			Organizer:   request.Organizer,
			Location:    null.StringFromPtr(request.Location),
			Latitude:    null.FloatFromPtr(request.Latitude),
			Longitude:   null.FloatFromPtr(request.Longitude),
			StartAt:     request.StartAt,
			EndAt:       request.EndAt,
			Status:      string(enum.EventStatusActive),
		}

		y1, m1, d1 := time.Time(request.StartAt).Date()
		y2, m2, d2 := time.Time(request.EndAt).Date()
		start := time.Date(y1, m1, d1, 0, 0, 0, 0, time.Time(request.StartAt).Location())
		end := time.Date(y2, m2, d2, 0, 0, 0, 0, time.Time(request.EndAt).Location())
		if start.Before(time.Now()) {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Start date must be after current date")
		}
		if end.Before(time.Now()) {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "End date must be after current date")
		}
		if end.Before(start) {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "End date must be after start date")
		}

		for _, ticket := range request.Tickets {
			event.Tickets = append(event.Tickets, &entity.Ticket{
				Category: ticket.Category,
				Price:    null.FloatFromPtr(ticket.Price).ValueOrZero(),
			})
		}

		// If no ticket, assume free event
		if len(event.Tickets) == 0 {
			event.Tickets = append(event.Tickets, &entity.Ticket{
				Category: "Free",
				Price:    0,
			})
		}

		if err := s.eventRepository.CreateEvent(ctx, tx, &event); err != nil {
			return err
		}

		return nil
	})
}

func (s *eventService) UpdateEvent(ctx context.Context, request dto.UpdateEventRequest) error {
	return s.eventRepository.WithTransaction(func(tx *gorm.DB) error {
		eventKey := "event:" + request.ID
		event, err := s.eventRepository.GetEventByID(ctx, tx, request.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Event not found")
		}

		if event.Status == string(enum.EventStatusFinished) {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Cannot update finished event")
		}

		isEnabled := null.BoolFromPtr(request.IsEnabled).ValueOrZero()
		// Created event can only update is_enabled
		if !isEnabled {
			event.Status = string(enum.EventStatusDisabled)
		} else {
			event.Status = string(enum.EventStatusActive)
		}

		if err = s.eventRepository.UpdateEvent(ctx, tx, event); err != nil {
			return err
		}

		if err = s.cache.Del(eventKey); err != nil {
			return err
		}

		return nil
	})
}

/* -------------------------------------------------------------------------- */
/*                                User Service                                */
/* -------------------------------------------------------------------------- */

func (s *eventService) GetUserEvents(ctx context.Context, queryParams url.Values) ([]entity.EventWithMinMaxPrice, *response.Meta, error) {
	var events []entity.EventWithMinMaxPrice
	var err error
	var meta *response.Meta
	db := s.eventRepository.SingleTransaction()

	db, meta = s.eventBuilder.ApplyBuilder(db, queryParams, &entity.EventWithMinMaxPrice{})
	events, err = s.eventRepository.GetActiveEventsWithMinMaxPrice(ctx, db)
	if err != nil {
		return nil, nil, err
	}

	return events, meta, nil
}

func (s *eventService) GetUserEventByID(ctx context.Context, id string) (*entity.Event, error) {
	event, err := s.GetEventByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if event.Status == string(enum.EventStatusDisabled) {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Event not found")
	}

	return event, nil
}

func (s *eventService) RegisterEvent(ctx context.Context, request dto.CreateEventRequest, userID string) error {
	return s.eventRepository.WithTransaction(func(tx *gorm.DB) error {
		user, err := s.userRepository.GetUserByID(ctx, tx, userID)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Invalid user")
		}

		eventApproval := entity.EventApproval{
			Status:      string(enum.EventApprovalStatusPending),
			Title:       request.Title,
			Description: null.StringFromPtr(request.Description),
			Organizer:   request.Organizer,
			Location:    null.StringFromPtr(request.Location),
			Latitude:    null.FloatFromPtr(request.Latitude),
			Longitude:   null.FloatFromPtr(request.Longitude),
			StartAt:     request.StartAt,
			EndAt:       request.EndAt,
			User:        user,
		}
		for _, ticket := range request.Tickets {
			eventApproval.EventApprovalsTickets = append(eventApproval.EventApprovalsTickets, &entity.EventApprovalTicket{
				Category: ticket.Category,
				Price:    null.FloatFromPtr(ticket.Price).ValueOrZero(),
			})
		}

		if err = s.eventApprovalRepository.CreateEventApproval(ctx, tx, &eventApproval); err != nil {
			return err
		}

		return nil
	})
}
