package redis

type Config struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
