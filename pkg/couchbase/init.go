package couchbase

import (
	"time"

	"github.com/couchbase/gocb/v2"
)

type Couchbase struct {
	Cluster *gocb.Cluster
}

func NewCouchbase(config *Config) (*Couchbase, error) {
	cluster, err := connect(config)
	if err != nil {
		return nil, err
	}

	return &Couchbase{Cluster: cluster}, nil
}

func connect(config *Config) (*gocb.Cluster, error) {
	cluster, err := gocb.Connect(
		config.Host,
		gocb.ClusterOptions{
			Username: config.Username,
			Password: config.Password,
		},
	)
	if err != nil {
		return nil, err
	}

	duration := time.Duration(config.WaitUntilReadySecs) * time.Second
	err = cluster.WaitUntilReady(duration, nil)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}
