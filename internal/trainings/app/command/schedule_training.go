package command

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/truesch/wild-workouts-go-ddd-example/internal/common/logs"
	"github.com/truesch/wild-workouts-go-ddd-example/internal/trainings/domain/training"
)

type ScheduleTraining struct {
	TrainingUUID string

	UserUUID string
	UserName string

	TrainingTime time.Time
	Notes        string
}

type ScheduleTrainingHandler struct {
	repo           training.Repository
	userService    UserService
	trainerService TrainerService
}

func NewScheduleTrainingHandler(repo training.Repository, userService UserService, trainerService TrainerService) ScheduleTrainingHandler {
	if repo == nil {
		panic("nil repo")
	}
	if userService == nil {
		panic("nil repo")
	}
	if trainerService == nil {
		panic("nil trainerService")
	}

	return ScheduleTrainingHandler{repo: repo, userService: userService, trainerService: trainerService}
}

func (h ScheduleTrainingHandler) Handle(ctx context.Context, cmd ScheduleTraining) (err error) {
	defer func() {
		logs.LogCommandExecution("ScheduleTraining", cmd, err)
	}()

	// TODO:
	topic, err := h.trainerService.GetTopic(ctx, cmd.TrainingTime)
	if err != nil {
		return errors.Wrap(err, "unable to schedule training")
	}
	log.Printf("Returned topic: %s", topic)

	tr, err := training.NewTraining(cmd.TrainingUUID, cmd.UserUUID, cmd.UserName, cmd.TrainingTime, topic)
	if err != nil {
		return err
	}

	if err := h.repo.AddTraining(ctx, tr); err != nil {
		return err
	}

	err = h.userService.UpdateTrainingBalance(ctx, tr.UserUUID(), -1)
	if err != nil {
		return errors.Wrap(err, "unable to change trainings balance")
	}

	err = h.trainerService.ScheduleTraining(ctx, tr.Time())
	if err != nil {
		return errors.Wrap(err, "unable to schedule training")
	}

	return nil
}
