package utils

import (
	"log/slog"
	"time"
)

type Timer struct {
	totalTimer  time.Time
	sectionTime time.Time
	timerName   string
}

func StartTimer(message string) Timer {
	slog.Info("Starting to time: " + message)
	return Timer{
		totalTimer:  time.Now(),
		sectionTime: time.Now(),
		timerName:   message,
	}
}

func (t *Timer) LogTime(message string) {
	slog.Info(
		message,
		"Timer",
		t.timerName,
		"Segement Elapsed Time",
		time.Duration(time.Since(t.sectionTime)).String(),
	)
	t.sectionTime = time.Now()
}

func (t *Timer) LogTotalTime() {
	slog.Info(t.timerName, "Total Elapsed Time", time.Duration(time.Since(t.totalTimer)).String())
}
