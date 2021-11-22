package command

import (
	"context"
	"time"

	"github.com/truesch/wild-workouts-go-ddd-example/internal/common/errors"
	"github.com/truesch/wild-workouts-go-ddd-example/internal/trainer/domain/hour"
)

type MakeHoursAvailableHandler struct {
	hourRepo hour.Repository
}

func NewMakeHoursAvailableHandler(hourRepo hour.Repository) MakeHoursAvailableHandler {
	if hourRepo == nil {
		panic("hourRepo is nil")
	}

	return MakeHoursAvailableHandler{hourRepo: hourRepo}
}

func (c MakeHoursAvailableHandler) Handle(ctx context.Context, hours []time.Time, topics []string, tags []string) error {
	for i, hourToUpdate := range hours {
		if err := c.hourRepo.UpdateHour(ctx, hourToUpdate, func(h *hour.Hour) (*hour.Hour, error) {
			if err := h.MakeAvailable(topics[i], tags[i]); err != nil {
				return nil, err
			}
			return h, nil
		}); err != nil {
			return errors.NewSlugError(err.Error(), "unable-to-update-availability")
		}
	}

	return nil
}
