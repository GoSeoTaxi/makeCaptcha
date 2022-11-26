package preGenerator

import (
	"go.uber.org/zap"
	"main/internal/config"
	"main/internal/models"
	"strconv"
	"sync"
	"time"
)

const itemsChanel = 10 //Размер канала
// const itemsGenerator = 100000  //Размер БД
const timeSleepTOStart = 1     //ТаймАут перед стартом
const timeOutGenerator = 10000 //NanoSec ожидание перед генерацией.

func StartGenerator(logger *zap.Logger, cfg *config.Config) (data chan models.SendData) {

	//create Storage
	mapS := sync.Map{}
	go generator(&mapS, logger, cfg)

	time.Sleep(timeSleepTOStart * time.Second)

	//Create queue
	c := make(chan models.SendData, itemsChanel)
	go readerMapTOChanel(c, &mapS)

	time.Sleep(timeSleepTOStart * time.Second)
	return c
}

func readerMapTOChanel(c chan models.SendData, m *sync.Map) {
	for {
		m.Range(func(key interface{}, value interface{}) bool {
			c <- value.(models.SendData)
			return true
		})
	}
}

func generator(m *sync.Map, logger *zap.Logger, cfg *config.Config) {

	items, err := strconv.Atoi(cfg.SizeCache)
	if err != nil {
		logger.Fatal("ERROR SIZE CACHE ", zap.Error(err))
	}

	for {
		logger.Info(`START GENERATOR`)
		for i := 0; i < items; i++ {
			passWordString, pic, err := models.CreatorCaptcha()
			if err != nil {
				logger.Error("no make Captcha")
				continue
			}

			m.Store(i, models.SendData{
				PassWord: passWordString,
				Picture:  pic,
			})
			time.Sleep(timeOutGenerator * time.Nanosecond)

			//	t := time.Now().Second()
			/*		if (time.Now().Second() % 20) == 0 {
					fmt.Println(i)
				}*/

		}
		logger.Info(`STOP GENERATOR`)
	}
	//	defer logger.Fatal(`err generator`)
}

/*
func generator_old(c chan models.SendData, logger *zap.Logger) {

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
}*/
