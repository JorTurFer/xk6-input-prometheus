package prometheus_test

import (
	"testing"

	prometheus "github.com/JorTurFer/xk6-input-prometheus"
	"github.com/stretchr/testify/assert"
)

func TestQueryScalar(t *testing.T) {
	module := prometheus.Prometheus{}
	client := module.NewPrometheusClient("http://demo.robustperception.io:9090", "", "")

	result, err := client.QueryScalar("up")
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, result, int64(1))
}

func TestQueryScalarWithOperation(t *testing.T) {
	module := prometheus.Prometheus{}
	client := module.NewPrometheusClient("http://demo.robustperception.io:9090", "", "")

	result, err := client.QueryScalar("sum(alertmanager_notifications_total)")
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, result, int64(1))
}

func TestQueryScalarWithBasicAuth(t *testing.T) {
	module := prometheus.Prometheus{}
	client := module.NewPrometheusClient("http://demo.robustperception.io:9090", "test", "1234")

	result, err := client.QueryScalar("sum(alertmanager_notifications_total)")
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, result, int64(1))
}

func TestQueryScalarWithError(t *testing.T) {
	module := prometheus.Prometheus{}
	client := module.NewPrometheusClient("http://demo.robustperception.io:9090", "test", "1234")

	_, err := client.QueryScalar("sum(random_metric)")
	assert.Error(t, err)
}
