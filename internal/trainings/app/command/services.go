package command

import (
	"context"
	"time"
)

type UserService interface {
	UpdateTrainingBalance(ctx context.Context, userID string, amountChange int) error
}

type TrainerService interface {
	ScheduleTraining(ctx context.Context, trainingTime time.Time) error
	CancelTraining(ctx context.Context, trainingTime time.Time) error
	GetTopic(ctx context.Context, trainingTime time.Time) (string, error)

	MoveTraining(
		ctx context.Context,
		newTime time.Time,
		originalTrainingTime time.Time,
	) error
}
