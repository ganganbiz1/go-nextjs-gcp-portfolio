package datadog

import (
	"context"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type IfDatadogClient interface {
	ServiceName(resource string) string
	StartTrace()
	StopTrace()
	StartSpan(ctx context.Context, resource string) (tracer.Span, context.Context)
}
