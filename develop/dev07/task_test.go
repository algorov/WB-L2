package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	// Начало отсчёта.
	now := time.Now()

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	// Объединение done-каналов.
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(7*time.Second),
	)

	// Фиксация времени после получения результата.
	after := time.Since(now)

	// Коэффициент расхождения полученного времени от ожидаемого.
	delta := 100 * (1 - int(after)/int(7*time.Second))

	// assert.
	assert.Equal(t, 0, delta)
}
