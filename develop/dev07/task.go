package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func main() {
	start := time.Now()

	// Возвращает канал, который позже через некоторое количество времени закроется посредством другой горутинки-слушателя.
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	// Ждёт завершения результирующего канала.
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(10*time.Second),
	)

	fmt.Printf("fone after %v", time.Since(start))
}

// Принимает неопределенное количество каналов и возвращает один канал, который свидетельствует, что все переданные каналы - done.
func or(channels ...<-chan interface{}) <-chan interface{} {
	// Если остался один канал, то просто его же возвращает, канал сам закроется после некоторого времени.
	if len(channels) == 1 {
		return channels[0]
	}

	// Результирующий канал, который возвращается.
	result := make(chan interface{})

	// Горутина, которая обрабатывает закрытые каналы, после обработки, закрывает результирующий канал.
	go func() {
		defer close(result)
		// Ждёт завершения первого канала.
		<-channels[0]

		// Ждёт завершения других каналов через рекурсию.
		<-or(channels[1:]...)
	}()

	// Возвращает результирующий канал.
	return result
}
