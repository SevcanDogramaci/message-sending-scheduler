package scheduler

import (
	"log"
	"time"

	"github.com/SevcanDogramaci/message-sending-scheduler/config"
)

type SchedulerMessageService interface {
	SendMessages() error
}

type Scheduler struct {
	config         *config.SchedulerConfig
	ticker         *time.Ticker
	done           chan bool
	messageService SchedulerMessageService
}

func NewScheduler(config *config.SchedulerConfig, service SchedulerMessageService) *Scheduler {
	return &Scheduler{
		config:         config,
		messageService: service,
	}
}

func (s *Scheduler) IsStarted() bool {
	return s.ticker != nil
}

func (s *Scheduler) IsDone() bool {
	_, ok := <-s.done
	return ok
}

func (s *Scheduler) Start() error {
	if s.IsStarted() {
		log.Println("A scheduler is already running")
		return nil
	}

	duration := time.Duration(s.config.PeriodSecs) * time.Second
	ticker := time.NewTicker(duration)
	done := make(chan bool)

	s.ticker = ticker
	s.done = done

	go func() {
		for {
			select {
			case <-done:
				log.Println("Scheduler got stop signal, stopping...")
				return
			case t := <-ticker.C:
				log.Println("Scheduler tick at", t)

				err := s.messageService.SendMessages()
				if err != nil {
					log.Println("Error occurred. Err:", err.Error())
				}
			}
		}
	}()

	log.Println("Scheduler started")
	return nil
}

func (s *Scheduler) Stop() error {
	if s.ticker != nil {
		s.ticker.Stop()
	}
	s.ticker = nil

	if s.done != nil {
		s.done <- true
		close(s.done)
	}

	log.Println("Scheduler stopped")
	return nil
}
