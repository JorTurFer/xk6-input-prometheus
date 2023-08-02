package prometheus_test

import (
	"testing"
	"time"

	prometheus "github.com/JorTurFer/xk6-input-prometheus"
	"github.com/stretchr/testify/assert"
)

const template string = "2006-01-02T15:04:05Z"

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

func TestQueryWithEmptyResponse(t *testing.T) {
	module := prometheus.Prometheus{}
	client := module.NewPrometheusClient("http://demo.robustperception.io:9090", "", "")

	response, err := client.Query("not-existing-metric")
	assert.NoError(t, err)
	assert.Empty(t, response.String())
}

func TestQueryWithOperationFail(t *testing.T) {
	module := prometheus.Prometheus{}
	client := module.NewPrometheusClient("http://demo.robustperception.io:9090", "", "")

	response, err := client.Query("rate(not-existing[5m])")
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestQueryRangeWithOperation(t *testing.T) {
	module := prometheus.Prometheus{}
	client := module.NewPrometheusClient("http://demo.robustperception.io:9090", "", "")
	start := time.Now().Add(-5 * time.Minute).UTC().Format(template)
	end := time.Now().UTC().Format(template)

	response, err := client.QueryRange("rate(prometheus_tsdb_head_samples_appended_total[5m])", start, end, "minute")
	assert.NoError(t, err)
	assert.NotEmpty(t, response.String())
}

func TestQueryRangeWithBasicAuth(t *testing.T) {
	module := prometheus.Prometheus{}
	client := module.NewPrometheusClient("http://demo.robustperception.io:9090", "test", "1234")
	start := time.Now().Add(-5 * time.Minute).UTC().Format(template)
	end := time.Now().UTC().Format(template)

	response, err := client.QueryRange("rate(prometheus_tsdb_head_samples_appended_total[5m])", start, end, "minute")
	assert.NoError(t, err)
	assert.NotEmpty(t, response.String())
}

func TestQueryRangeWithOperationFail(t *testing.T) {
	module := prometheus.Prometheus{}
	client := module.NewPrometheusClient("http://demo.robustperception.io:9090", "", "")
	start := time.Now().Add(-5 * time.Minute).UTC().Format(template)
	end := time.Now().UTC().Format(template)

	response, err := client.QueryRange("rate(not-existing[5m])", start, end, "minute")
	assert.Error(t, err)
	assert.Nil(t, response)
}
