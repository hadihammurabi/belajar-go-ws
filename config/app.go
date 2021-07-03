package config

import (
	"fmt"
	"os"
)

// AppConfig struct
type AppConfig struct {
	Port   string
	WsPort string
}

// ConfigureApp func
func ConfigureApp() *AppConfig {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	wsPort := os.Getenv("APP_WS_PORT")
	if wsPort == "" {
		wsPort = "8081"
	}

	return &AppConfig{
		Port:   fmt.Sprintf(":%s", port),
		WsPort: fmt.Sprintf(":%s", wsPort),
	}
}
