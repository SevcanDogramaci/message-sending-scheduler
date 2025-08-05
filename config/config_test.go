package config_test

import (
	"testing"

	"github.com/SevcanDogramaci/message-sending-scheduler/config"
	"github.com/SevcanDogramaci/message-sending-scheduler/pkg/couchbase"
	"github.com/SevcanDogramaci/message-sending-scheduler/pkg/redis"
	"github.com/stretchr/testify/assert"
)

func TestInitConfig_GivenEnv_ThenItShouldReadConfig(t *testing.T) {
	env := "test"
	expectedConfig := &config.AppConfig{
		Scheduler: &config.SchedulerConfig{
			PeriodSecs: 10,
		},
		Couchbase: &couchbase.Config{
			Host:               "test:8091",
			Username:           "test-username",
			Password:           "test-password",
			WaitUntilReadySecs: 5,
		},
		Webhook: &config.ClientConfig{
			URL:    "https://test.webhook.site/test-token",
			APIKey: "test-api-key",
		},
		Redis: &redis.Config{
			Host:           "test:6379",
			Password:       "test-password",
			DB:             1,
			DefaultTTLSecs: 10,
		},
	}
	actualConfig, err := config.InitConfigs(env)

	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, actualConfig)
}

func TestInitConfig_GivenBadConfig_ThenItShouldReturnError(t *testing.T) {
	env := "test.error"
	config, err := config.InitConfigs(env)

	assert.Error(t, err)
	assert.Nil(t, config)
}
