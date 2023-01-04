package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ktoshiya/golang-todo/clock"
	"github.com/ktoshiya/golang-todo/config"
	"github.com/ktoshiya/golang-todo/handler"
	"github.com/ktoshiya/golang-todo/service"
	"github.com/ktoshiya/golang-todo/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{status: "OK"}`))
	})

	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	clocker := clock.RealClocker{}
	r := store.Repository{Clocker: clocker}

	at := &handler.AddTask{Service: &service.AddTask{
		DB:   db,
		Repo: &r,
	}, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)

	lt := &handler.ListTask{Service: &service.ListTask{
		DB:   db,
		Repo: &r,
	}}
	mux.Get("/tasks", lt.ServeHTTP)

	ru := &handler.RegisterUser{Service: &service.RegisterUser{DB: db, Repo: &r}, Validator: v}
	mux.Post("/register", ru.ServeHTTP)

	return mux, cleanup, nil
}
