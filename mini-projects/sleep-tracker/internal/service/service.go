package service

import (
	"errors"
	"sync"
	"time"
)

type SleepService struct {
	mu         sync.RWMutex
	SonnikList []Sleep
	nextId     int
}

var (
	ErrorIncorrectDate = errors.New("incorrect date")
	ErrorIncorrectTime = errors.New("incorrect time")
)

var dateFormats = []string{
	"02.01.2006 15:04",
	"02 01 2006 15:04",
	"02-01-2006 15:04",
	"02/01/2006 15:04",
	"2006-01-02 15:04",
}

type UserSleep struct {
	SleepDate string `json:"sleep_date"`
	Bedtime   string `json:"bed_time"`
	WakeTime  string `json:"wake_time"`
	Quality   int    `json:"quality"`
	Mood      int    `json:"mood"`
	Notes     string `json:"notes"`
}

type Sleep struct {
	ID        int           `json:"id"`
	SleepDate time.Time     `json:"sleep_date"`
	Bedtime   time.Time     `json:"bed_time"`
	WakeTime  time.Time     `json:"wake_time"`
	Duration  time.Duration `json:"duration"`
	Quality   int           `json:"quality"`
	Mood      int           `json:"mood"`
	Notes     string        `json:"notes"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (s *SleepService) List() []Sleep {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]Sleep, len(s.SonnikList))
	copy(out, s.SonnikList)
	return out
}

func parseSleep(sleep UserSleep) (Sleep, error) {
	btime := sleep.Bedtime
}

func parseTimeFlexible(s string) (time.Time, error) {
	for _, f := range dateFormats {
		if t, err := time.Parse(f, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("unsupported time format")
}

func (s *SleepService) Add(sleep Sleep) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if sleep.SleepDate.IsZero() {
		return ErrorIncorrectDate
	}

	if sleep.Bedtime.IsZero() {
		return ErrorIncorrectTime
	}

	duration := time.Now().Sub(sleep.Bedtime)
	sleep.Duration = duration
	sleep.CreatedAt = time.Now()
}
