package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/truesch/wild-workouts-go-ddd-example/internal/common/logs"
	"github.com/truesch/wild-workouts-go-ddd-example/internal/common/server"
	"github.com/truesch/wild-workouts-go-ddd-example/internal/trainings/ports"
	"github.com/truesch/wild-workouts-go-ddd-example/internal/trainings/service"
)

func main() {
	logs.Init()

	ctx := context.Background()

	app, cleanup := service.NewApplication(ctx)
	defer cleanup()

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(app), router)
	})
}
