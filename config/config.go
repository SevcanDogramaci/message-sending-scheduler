package config

import "github.com/SevcanDogramaci/message-sending-scheduler/pkg/couchbase"

type AppConfig struct {
	Scheduler *SchedulerConfig  `json:"scheduler"`
	Couchbase *couchbase.Config `json:"couchbase"`
	Webhook   *ClientConfig     `json:"webhook"`
}

type ClientConfig struct {
	URL    string `json:"url"`
	APIKey string `json:"api_key"`
}

type SchedulerConfig struct {
	PeriodSecs int `json:"period_secs"`
}
