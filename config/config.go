package config

import (
	"github.com/SevcanDogramaci/message-sending-scheduler/pkg/couchbase"
	"github.com/SevcanDogramaci/message-sending-scheduler/pkg/redis"
)

type AppConfig struct {
	Scheduler *SchedulerConfig  `json:"scheduler"`
	Couchbase *couchbase.Config `json:"couchbase"`
	Webhook   *ClientConfig     `json:"webhook"`
	Redis     *redis.Config		`json:"redis"`
}

type ClientConfig struct {
	URL    string `json:"url"`
	APIKey string `json:"api_key"`
}

type SchedulerConfig struct {
	PeriodSecs int `json:"period_secs"`
}
