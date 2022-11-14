package handlers

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"main/internal/models"
	"net/http"
)

type Handler struct {
	logger *zap.Logger
	ctx    context.Context
}

func NewHandler(ctx context.Context, logger *zap.Logger) Handler {
	return Handler{
		logger: logger,
		ctx:    ctx,
	}
}

func (h *Handler) HandlerGetCaptcha() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		h.logger.Debug("getting captcha")
		passWordString, pic, err := models.CreatorCaptcha()
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		json, err := json.Marshal(models.SendData{PassWord: passWordString, Picture: pic})
		if err != nil {
			h.logger.Error("no make json")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)

		h.logger.Debug("sending captcha")
	}
}
