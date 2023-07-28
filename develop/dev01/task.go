package main

import (
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func GetTime(address string) (*time.Time, error) {
	ntpTime, err := ntp.Time(address)
	if err != nil {
		return nil, err
	}
	return &ntpTime, nil
}

func main() {
	// Получаем точное время с использованием NTP.
	address := "0.beevik-ntp.pool.ntp.org"
	ntpTime, err := GetTime(address)
	if err != nil {
		log.Println("error when getting time from NTP:", err)
		os.Exit(1)
	}
	// Получаем текущее локальное время.
	localTime := time.Now()
	// Выводим текущее время и точное время.
	log.Println("Текущее время:", localTime.String())
	log.Println("Точное время: ", ntpTime.String())
}
