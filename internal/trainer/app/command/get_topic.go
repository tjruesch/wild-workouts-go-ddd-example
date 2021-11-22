package command

import (
	"context"
	"log"
	"time"

	"github.com/truesch/wild-workouts-go-ddd-example/internal/common/errors"
	"github.com/truesch/wild-workouts-go-ddd-example/internal/trainer/domain/hour"
)

type GetTopicHandler struct {
	hourRepo hour.Repository
}

func NewGetTopicHandler(hourRepo hour.Repository) GetTopicHandler {
	if hourRepo == nil {
		panic("hourRepo is nil")
	}

	return GetTopicHandler{hourRepo: hourRepo}
}

func (c GetTopicHandler) Handle(ctx context.Context, hour time.Time) (string, error) {

	h, err := c.hourRepo.GetHour(ctx, hour)

	// FIXME:
	log.Printf("Hour: %#v, Topic: %s", h, h.Topic())

	if err != nil {
		return "", errors.NewSlugError(err.Error(), "unable-to-get-topic")
	}

	return h.Topic(), nil
}
