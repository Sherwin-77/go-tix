package service

import (
	"context"
	"encoding/json"
	"github.com/sherwin-77/go-tix/internal/entity"
	"github.com/sherwin-77/go-tix/internal/repository"
	"github.com/sherwin-77/go-tix/pkg/caches"
	"github.com/sherwin-77/go-tix/pkg/query"
	"github.com/sherwin-77/go-tix/pkg/response"
	"net/url"
	"time"
)

type EventService interface {
	GetEvents(ctx context.Context, queryParams url.Values) ([]entity.EventWithMinMaxPrice, *response.Meta, error)
	GetEventByID(ctx context.Context, id string) (*entity.Event, error)
}

type eventService struct {
	eventRepository repository.EventRepository
	eventBuilder    query.Builder
	cache           caches.Cache
}

func NewEventService(
	eventRepository repository.EventRepository,
	eventBuilder query.Builder,
	cache caches.Cache,
) EventService {
	return &eventService{eventRepository, eventBuilder, cache}
}

/* -------------------------------------------------------------------------- */
/*                            User & Event Service                            */
/* -------------------------------------------------------------------------- */

func (s *eventService) GetEvents(ctx context.Context, queryParams url.Values) ([]entity.EventWithMinMaxPrice, *response.Meta, error) {
	var events []entity.EventWithMinMaxPrice
	var err error
	var meta *response.Meta
	db := s.eventRepository.SingleTransaction()

	db, meta = s.eventBuilder.ApplyBuilder(db, queryParams, &entity.EventWithMinMaxPrice{})
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
			return nil, err
		}

		data, _ := json.Marshal(event)

		if err = s.cache.Set(eventKey, string(data), 5*time.Minute); err != nil {
			return nil, err
		}
	}

	return event, nil
}
