package ports

import (
	"context"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/truesch/wild-workouts-go-ddd-example/internal/common/genproto/trainer"
	"github.com/truesch/wild-workouts-go-ddd-example/internal/trainer/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) MakeHourAvailable(ctx context.Context, request *trainer.UpdateHourRequest) (*empty.Empty, error) {
	trainingTime, err := protoTimestampToTime(request.Time)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "unable to parse time")
	}

	if err := g.app.Commands.MakeHoursAvailable.Handle(
		ctx,
		[]time.Time{trainingTime},
		[]string{request.Topic},
		[]string{request.Tags},
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) ScheduleTraining(ctx context.Context, request *trainer.UpdateHourRequest) (*empty.Empty, error) {
	trainingTime, err := protoTimestampToTime(request.Time)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "unable to parse time")
	}

	if err := g.app.Commands.ScheduleTraining.Handle(ctx, trainingTime); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) CancelTraining(ctx context.Context, request *trainer.UpdateHourRequest) (*empty.Empty, error) {
	trainingTime, err := protoTimestampToTime(request.Time)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "unable to parse time")
	}

	if err := g.app.Commands.CancelTraining.Handle(ctx, trainingTime); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) IsHourAvailable(ctx context.Context, request *trainer.IsHourAvailableRequest) (*trainer.IsHourAvailableResponse, error) {
	trainingTime, err := protoTimestampToTime(request.Time)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "unable to parse time")
	}

	isAvailable, err := g.app.Queries.HourAvailability.Handle(ctx, trainingTime)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &trainer.IsHourAvailableResponse{IsAvailable: isAvailable}, nil
}

func (g GrpcServer) GetTopic(ctx context.Context, request *trainer.GetTopicRequest) (*trainer.GetTopicResponse, error) {
	trainingTime, err := protoTimestampToTime(request.Time)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "unable to parse time")
	}

	topic, err := g.app.Commands.GetTopic.Handle(ctx, trainingTime)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &trainer.GetTopicResponse{Topic: topic}, nil
}

func protoTimestampToTime(timestamp *timestamp.Timestamp) (time.Time, error) {
	t, err := ptypes.Timestamp(timestamp)
	if err != nil {
		return time.Time{}, errors.New("unable to parse time")
	}

	t = t.UTC().Truncate(time.Hour)

	return t, nil
}
