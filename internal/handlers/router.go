package handlers

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"main/internal/config"
	"main/internal/models"
	"strconv"
	"time"
)

//const timeOutConst = 5

func CaptchaRouter(ctx context.Context, c chan models.SendData, cfg *config.Config, logger *zap.Logger) chi.Router {

	r := chi.NewRouter()
	mh := NewHandler(ctx, c, logger)

	tou, err := strconv.Atoi(cfg.TimeOut500)
	if err != nil {
		logger.Fatal("ERROR read TimeOut ", zap.Error(err))
	}
	//	r.Use(middleware.RequestID)
	//	r.Use(middleware.RealIP)
	//r.Use(middleware.Logger)
	//	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Duration(tou) * time.Second))

	r.Route("/", func(r chi.Router) {
		r.Get("/get", Conveyor(mh.HandlerGetCaptcha(), packGZIP))
	})

	return r
}
