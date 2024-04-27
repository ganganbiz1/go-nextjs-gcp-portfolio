package newrelic

import (
	"context"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/config"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/logger"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Client struct {
	App *newrelic.Application
}

func NewClient(
	c *config.NewRelicConfig,
) *Client {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(c.AppName),
		newrelic.ConfigLicense(c.LicenseKey),
		newrelic.ConfigDistributedTracerEnabled(c.Enabled),
		newrelic.ConfigAppLogForwardingEnabled(c.LogForwardingEnabled),
	)

	if err != nil {
		logger.Warn("newrelic Initialization failure")
		return nil
	}

	return &Client{
		App: app,
	}
}

func (c *Client) StartTransaction(name string) {
	txn := c.App.StartTransaction(name)
	defer txn.End()
}

func (c *Client) NewContext(ctx context.Context, txn *newrelic.Transaction) context.Context {
	return newrelic.NewContext(ctx, txn)
}
