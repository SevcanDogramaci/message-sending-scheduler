package couchbase

type Config struct {
	Host               string `json:"host"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	WaitUntilReadySecs int    `json:"wait_until_ready_secs"`
}
