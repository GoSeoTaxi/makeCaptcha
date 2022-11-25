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
	chanel chan models.SendData
	ctx    context.Context
}

func NewHandler(ctx context.Context, c chan models.SendData, logger *zap.Logger) Handler {
	return Handler{
		logger: logger,
		chanel: c,
		ctx:    ctx,
	}
}

func (h *Handler) HandlerGetCaptcha() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		h.logger.Debug("getting captcha")

		data := <-h.chanel
		json, err := json.Marshal(models.SendData{PassWord: data.PassWord, Picture: data.Picture})
		if err != nil {
			h.logger.Error("no make json")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)

		h.logger.Debug("sending captcha")
	}
}
