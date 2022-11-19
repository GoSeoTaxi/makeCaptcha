package preGenerator

import (
	"go.uber.org/zap"
	"main/internal/models"
	"time"
)

const itemsGenerator = 3
const makingGenerator = 2
const timeSleepTOStart = 1

func StartGenerator(logger *zap.Logger) (data chan models.SendData) {

	c := make(chan models.SendData, itemsGenerator)

	go func(c chan models.SendData, logger *zap.Logger) {
		for i := 0; i < makingGenerator; i++ {
			go generator(c, logger)
		}
	}(c, logger)

	time.Sleep(timeSleepTOStart * time.Second)
	return c
}

func generator(c chan models.SendData, logger *zap.Logger) {

	for {
		passWordString, pic, err := models.CreatorCaptcha()
		if err != nil {
			logger.Error("no make Captcha")
			continue
		}

		c <- models.SendData{
			PassWord: passWordString,
			Picture:  pic,
		}
	}
}
