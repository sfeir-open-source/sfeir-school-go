package statistics

import (
	logger "github.com/sirupsen/logrus"
	"time"
)

const (
	statisticsChannelSize = 1000
)

// Statistics is the worker to persist the request statistics
type Statistics struct {
	statistics    chan bool
	counter       uint32
	start         time.Time
	loggingPeriod time.Duration
}

// NewStatistics creates a new statistics structure and launches its worker routine
func NewStatistics(loggingPeriod time.Duration) *Statistics {
	sw := Statistics{
		statistics:    make(chan bool, statisticsChannelSize),
		counter:       0,
		start:         time.Now(),
		loggingPeriod: loggingPeriod,
	}
	go sw.run()
	return &sw
}

// PlusOne is used to send a statistics hit increment
func (s *Statistics) PlusOne() {
	s.statistics <- true
}

func (s *Statistics) run() {
	ticker := time.NewTicker(s.loggingPeriod)
	for {
		select {
		case stat := <-s.statistics:
			logger.WithField("stat", stat).Debug("new count received")
			s.counter++
		case <-ticker.C:
			elapsed := time.Since(s.start)
			rate := float64(s.counter) / elapsed.Seconds()
			logger.WithField("request_rate", rate).Warn("request monitoring")
			s.counter = 0
			s.start = time.Now()
		}
	}
}
