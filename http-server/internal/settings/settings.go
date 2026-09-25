package settings

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	settingsPath = "settings"
	settingsName = "settings.json"
)

type Settings struct {
	Server HttpServer `json:"server"`
}

type HttpServer struct {
	Host 					string 	`json:"host"`
	Port 					int    	`json:"port"`
	ShutdownTimeoutSecond 	int 	`json:"timeout_shutdown_second"`
}

func Load(env string) (*Settings, error) {
	var set Settings
	path := fmt.Sprintf("%s/%s_%s", settingsPath, env, settingsName)

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Load file error: %w", err)
	}

	if err := json.Unmarshal(data, &set); err != nil {
		return nil, fmt.Errorf("Unmarshan erro: %w", err)
	}

	return &set, nil
}