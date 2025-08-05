package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func InitConfigs(env string) (*AppConfig, error) {
	filePath := fmt.Sprintf("./.configs/%s.json", env)
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var config AppConfig
	if err = json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
