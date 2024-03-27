package main

import (
	"testing"
	"time"
)

func TestM(t *testing.T) {
	start := time.Now()

	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(14*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
		sig(2*time.Second),
	)

	duration := time.Since(start)
	expected := 2 * time.Second // Измените ожидаемое значение на то, которое вам нужно

	if duration < expected {
		t.Errorf("Expected duration to be greater than or equal to %v, but got %v", expected, duration)
	}
}
