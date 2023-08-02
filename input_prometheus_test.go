package prometheus_test

import (
	"testing"

	prometheus "github.com/JorTurFer/xk6-input-prometheus"
	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	module := prometheus.Prometheus{}
	client := module.NewPrometheusClient("http://demo.robustperception.io:9090", "", "")

	response, err := client.Query("up")
	assert.NoError(t, err)
	assert.NotEmpty(t, response.String())
}

func TestQueryWithOperation(t *testing.T) {
	module := prometheus.Prometheus{}
	client := module.NewPrometheusClient("http://demo.robustperception.io:9090", "", "")

	response, err := client.Query("sum(alertmanager_notifications_total)")
	assert.NoError(t, err)
	assert.NotEmpty(t, response.String())
}

func TestQueryWithBasicAuth(t *testing.T) {
	module := prometheus.Prometheus{}
	client := module.NewPrometheusClient("http://demo.robustperception.io:9090", "test", "1234")

	response, err := client.Query("sum(alertmanager_notifications_total)")
	assert.NoError(t, err)
	assert.NotEmpty(t, response.String())
}

func TestQueryScalarWithError(t *testing.T) {
	module := prometheus.Prometheus{}
	client := module.NewPrometheusClient("http://demo.robustperception.io:9090", "test", "1234")

	response, err := client.Query("random_metrics")
	assert.NoError(t, err)
	assert.Empty(t, response.String())
}
