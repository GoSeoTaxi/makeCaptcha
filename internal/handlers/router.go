package handlers

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// BonusRouter arranges the whole API endpoints and their correponding handlers
func CaptchaRouter(ctx context.Context, logger *zap.Logger) chi.Router {

	r := chi.NewRouter()
	mh := NewHandler(ctx, logger)

	//	r.Use(middleware.RequestID)
	//	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	//	r.Use(middleware.Recoverer)

	r.Route("/", func(r chi.Router) {
		r.Get("/get", Conveyor(mh.HandlerGetCaptcha(), packGZIP))
	})

	return r
}
