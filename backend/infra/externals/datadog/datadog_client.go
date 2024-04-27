package datadog

import (
	"context"
	"fmt"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/config"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/datadog"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type Client struct {
	config *config.DatadogConfig
}

func NewClient(config *config.DatadogConfig) datadog.IfDatadogClient {
	return &Client{config}
}

func (c *Client) ServiceName(resource string) string {
	return c.config.ServiceNamePrefix + resource
}
func (c *Client) StartTrace() {
	tracer.Start(
		tracer.WithAgentAddr(fmt.Sprintf("%s:%s", c.config.AgentHost, c.config.AgentPort)),
		tracer.WithService(c.ServiceName("api")),
		tracer.WithEnv(c.config.ENV))
}
func (c *Client) StopTrace() {
	tracer.Stop()
}
func (c *Client) StartSpan(ctx context.Context, resource string) (tracer.Span, context.Context) {
	return tracer.StartSpanFromContext(ctx, "parent.request",
		tracer.SpanType(ext.SpanTypeWeb),
		tracer.ServiceName(c.config.ServiceNamePrefix+resource),
		tracer.ResourceName("/"+resource),
	)
}
