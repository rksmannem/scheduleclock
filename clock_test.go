package main

import (
	"testing"
	"time"
)

func TestClockForFiftySeconds(t *testing.T) {
	Clock(time.Duration(50) * time.Second)
}
func TestClockForOneMinute(t *testing.T) {
	Clock(time.Duration(1) * time.Minute)
}

func TestClockForFiveMinute(t *testing.T) {
	Clock(time.Duration(5) * time.Minute)
}

func TestClockForTwoHours(t *testing.T) {
	Clock(time.Duration(2) * time.Hour)
}
