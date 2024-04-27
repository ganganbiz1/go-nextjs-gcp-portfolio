package config

import (
	"os"
	"strconv"
)

type NewRelicConfig struct {
	AppName                  string
	LicenseKey               string
	Enabled                  bool
	DistributedTracerEnabled bool
	LogForwardingEnabled     bool
}

func NewNewRelicConfig() *NewRelicConfig {
	enabled, err := strconv.ParseBool(os.Getenv("NEW_RELIC_ENABLED"))
	if err != nil {
		enabled = false
	}

	traceEnabled, err := strconv.ParseBool(os.Getenv("NEW_RELIC_TRACE_ENABLED"))
	if err != nil {
		traceEnabled = false
	}

	logForwardingEnabled, err := strconv.ParseBool(os.Getenv("NEW_RELIC_LOG_FORWARDING_ENABLED"))
	if err != nil {
		traceEnabled = false
	}

	return &NewRelicConfig{
		AppName:                  os.Getenv("APP_NAME"),
		LicenseKey:               os.Getenv("NEW_RELIC_LICENSE_KEY"),
		Enabled:                  enabled,
		DistributedTracerEnabled: traceEnabled,
		LogForwardingEnabled:     logForwardingEnabled,
	}
}
