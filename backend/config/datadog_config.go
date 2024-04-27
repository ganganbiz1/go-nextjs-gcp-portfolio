package config

import (
	"os"
)

type DatadogConfig struct {
	AgentHost         string
	AgentPort         string
	ServiceNamePrefix string
	ENV               string
}

func NewDatadogConfig() *DatadogConfig {
	return &DatadogConfig{
		AgentHost:         os.Getenv("DD_AGENT_HOST"),
		AgentPort:         os.Getenv("DD_AGENT_PORT"),
		ServiceNamePrefix: os.Getenv("APP_ENV") + os.Getenv("APP_NAME") + "-",
		ENV:               os.Getenv("APP_ENV"),
	}
}
