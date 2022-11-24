package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"main/internal/config"
	"main/internal/handlers"
	"main/internal/models"
	"main/internal/preGenerator"
	"net/http"
	"strconv"
	"time"
)

func main() {

	fmt.Print("starting...")
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("can't load config: %v", err)
	}

	logger, err := config.InitLogger(cfg.Debug, cfg.AppName)
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	logger.Info("initializing the service...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c1 := preGenerator.StartGenerator(logger)
	//go printerChanel(c1, logger)

	// prepare handles
	r := handlers.CaptchaRouter(ctx, c1, logger)
	srv := &http.Server{Addr: cfg.Endpoint, Handler: r}

	logger.Info("Start serving on", zap.String("endpoint name", cfg.Endpoint))
	log.Fatal(srv.ListenAndServe())

}

func printerChanel(c chan models.SendData, logger *zap.Logger) {
	for {
		logger.Info("Queue", zap.String("queue - chan", time.Now().Format("2006-01-02 15:04:05")+" q="+strconv.Itoa(len(c))))
		time.Sleep(100 * time.Microsecond)
	}
}
