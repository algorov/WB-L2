package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	// Определяет время с сервера и хоста.
	ntpTime := GetTime()
	hostTime := time.Now()

	// Разница между временами.
	different := hostTime.Sub(hostTime)

	// Если разница отрицательна, делает положительным
	if different < 0 {
		different = -different
	}

	// Если разница более в одну секунду, то тест провален.
	flag := false
	if different > time.Second {
		flag = true
	}

	// Сравнивание ожидаемого результата с полученным.
	assert.Equal(t, false, flag,
		fmt.Sprintf("Time from NTP: %+v\nTIme from host: %+v\nDifferent: %d", ntpTime, hostTime, different))
}
