package main

import (
	"fmt"
	"reflect"
	"time"
)

// or объединяет несколько done-каналов в один.
var or = func(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	// Создаем результирующий канал
	result := make(chan interface{})

	go func() {
		defer close(result)

		// Используем select для обработки всех каналов
		var cases []reflect.SelectCase
		for _, ch := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			})
		}

		// Ожидаем закрытия любого канала
		_, _, _ = reflect.Select(cases)
	}()

	return result
}

func main() {
	// Функция для создания сигнального канала
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	// Используем or для объединения нескольких каналов
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v\n", time.Since(start))
}
