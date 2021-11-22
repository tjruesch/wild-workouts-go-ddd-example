package app

import (
	"github.com/truesch/wild-workouts-go-ddd-example/internal/trainer/app/command"
	"github.com/truesch/wild-workouts-go-ddd-example/internal/trainer/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CancelTraining   command.CancelTrainingHandler
	ScheduleTraining command.ScheduleTrainingHandler

	MakeHoursAvailable   command.MakeHoursAvailableHandler
	MakeHoursUnavailable command.MakeHoursUnavailableHandler

	GetTopic command.GetTopicHandler
}

type Queries struct {
	HourAvailability      query.HourAvailabilityHandler
	TrainerAvailableHours query.AvailableHoursHandler
}
