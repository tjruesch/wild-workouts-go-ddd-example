package adapters

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"

	"github.com/truesch/wild-workouts-go-ddd-example/internal/common/genproto/trainer"
)

type TrainerGrpc struct {
	client trainer.TrainerServiceClient
}

func NewTrainerGrpc(client trainer.TrainerServiceClient) TrainerGrpc {
	return TrainerGrpc{client: client}
}

func (s TrainerGrpc) ScheduleTraining(ctx context.Context, trainingTime time.Time) error {
	timestamp, err := ptypes.TimestampProto(trainingTime)
	if err != nil {
		return errors.Wrap(err, "unable to convert time to proto timestamp")
	}

	_, err = s.client.ScheduleTraining(ctx, &trainer.UpdateHourRequest{
		Time: timestamp,
	})

	return err
}

func (s TrainerGrpc) CancelTraining(ctx context.Context, trainingTime time.Time) error {
	timestamp, err := ptypes.TimestampProto(trainingTime)
	if err != nil {
		return errors.Wrap(err, "unable to convert time to proto timestamp")
	}

	_, err = s.client.CancelTraining(ctx, &trainer.UpdateHourRequest{
		Time: timestamp,
	})

	return err
}

func (s TrainerGrpc) MoveTraining(
	ctx context.Context,
	newTime time.Time,
	originalTrainingTime time.Time,
) error {
	err := s.ScheduleTraining(ctx, newTime)
	if err != nil {
		return errors.Wrap(err, "unable to schedule training")
	}

	err = s.CancelTraining(ctx, originalTrainingTime)
	if err != nil {
		return errors.Wrap(err, "unable to cancel training")
	}

	return nil
}

func (s TrainerGrpc) GetTopic(ctx context.Context, trainingTime time.Time) (string, error) {
	timestamp, err := ptypes.TimestampProto(trainingTime)
	if err != nil {
		return "", errors.Wrap(err, "unable to convert time to proto timestamp")
	}

	topicResponse, err := s.client.GetTopic(ctx, &trainer.GetTopicRequest{
		Time: timestamp,
	})

	log.Printf("Response: %v, Topic: %s", topicResponse, topicResponse.Topic)

	return topicResponse.GetTopic(), err
}
