package definitions

import (
	"time"

	"github.com/Axway/agents-gravitee/client/pkg/gravitee/models"
)

type StatsClient interface {
	GetEnvironments() []string
	GetStats(env, dimension, metricSelect string, start, end time.Time) (*models.Metrics, error)
	GetProduct(productName string) (*models.ApiProduct, error)
}
