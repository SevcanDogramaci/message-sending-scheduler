package scheduler_test

import (
	"testing"

	"github.com/SevcanDogramaci/message-sending-scheduler/config"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/mocks"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/scheduler"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestStart_ItShouldStartTicker(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockService := mocks.NewMockSchedulerMessageService(mockController)
	scheduler := scheduler.NewScheduler(
		&config.SchedulerConfig{PeriodSecs: 1}, mockService)

	assert.False(t, scheduler.IsStarted())

	scheduler.Start()

	assert.True(t, scheduler.IsStarted())
}


func TestStop_ItShouldStopTicker(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockService := mocks.NewMockSchedulerMessageService(mockController)
	scheduler := scheduler.NewScheduler(
		&config.SchedulerConfig{PeriodSecs: 1}, mockService)

	scheduler.Start()	
	scheduler.Stop()

	assert.False(t, scheduler.IsDone())
}
