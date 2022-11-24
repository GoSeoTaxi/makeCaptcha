package handlers

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"main/internal/models"
	"time"
)

const timeOutConst = 5

func CaptchaRouter(ctx context.Context, c chan models.SendData, logger *zap.Logger) chi.Router {

	r := chi.NewRouter()
	mh := NewHandler(ctx, c, logger)

	//	r.Use(middleware.RequestID)
	//	r.Use(middleware.RealIP)
	//r.Use(middleware.Logger)
	//	r.Use(middleware.Recoverer)
	//	r.Use(middleware.Timeout(timeOut * time.Second))
	r.Use(middleware.Timeout(timeOutConst * time.Second))

	r.Route("/", func(r chi.Router) {
		r.Get("/get", Conveyor(mh.HandlerGetCaptcha(), packGZIP))
	})

	return r
}
