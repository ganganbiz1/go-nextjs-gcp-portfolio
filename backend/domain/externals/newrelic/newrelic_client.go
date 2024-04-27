package newrelic

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type IfNewrelicClient interface {
	StartTransaction(name string)
	NewContext(ctx context.Context, txn *newrelic.Transaction)
}
