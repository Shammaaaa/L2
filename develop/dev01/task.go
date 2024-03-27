package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Установка соединения с NTP-сервером и получение точного времени
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		log.Printf("Ошибка при получении времени: %v", err)
		os.Exit(1)
	}

	// Печать текущего точного времени
	fmt.Println("Текущее точное время:", ntpTime.Format(time.RFC3339))
}
