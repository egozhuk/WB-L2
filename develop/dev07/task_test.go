package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	// Вспомогательная функция для создания сигнального канала
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	t.Run("Single channel", func(t *testing.T) {
		start := time.Now()
		<-or(sig(1 * time.Second))
		duration := time.Since(start)
		if duration < 1*time.Second || duration > 2*time.Second {
			t.Errorf("Expected ~1s, got %v", duration)
		}
	})

	t.Run("Multiple channels", func(t *testing.T) {
		start := time.Now()
		<-or(
			sig(2*time.Second),
			sig(1*time.Second),
			sig(3*time.Second),
		)
		duration := time.Since(start)
		if duration < 1*time.Second || duration > 2*time.Second {
			t.Errorf("Expected ~1s, got %v", duration)
		}
	})

	t.Run("No channels", func(t *testing.T) {
		ch := or()
		if ch != nil {
			t.Error("Expected nil for no channels")
		}
	})
}
