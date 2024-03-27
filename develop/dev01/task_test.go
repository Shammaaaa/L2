package main

import (
	"github.com/beevik/ntp"
	"testing"
)

func TestGetNTPTime_Success(t *testing.T) {
	ntpServer := "pool.ntp.org"
	ntpTime, err := ntp.Time(ntpServer)
	if err != nil {
		t.Errorf("Ошибка при получении времени: %v", err)
	}

	// Проверка, что время получено
	if ntpTime.IsZero() {
		t.Error("Время не было получено")
	}

}
