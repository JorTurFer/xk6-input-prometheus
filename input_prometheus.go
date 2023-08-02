package prometheus

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"go.k6.io/k6/js/modules"
)

// init is called by the Go runtime at application startup.
func init() {
	modules.Register("k6/x/prometheusread", new(Prometheus))
}

type Prometheus struct{}

type Client struct {
	url      string
	username string
	password config.Secret
}

func (*Prometheus) NewPrometheusClient(url, username string, password config.Secret) Client {
	return Client{
		url:      url,
		username: username,
		password: password,
	}
}

func (c *Client) Query(query string) (model.Value, error) {
	client, err := c.generateClient()
	if err != nil {
		return nil, err
	}
	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, warnings, err := v1api.Query(ctx, query, time.Now(), v1.WithTimeout(5*time.Second))
	if err != nil || len(warnings) > 0 {
		return nil, err
	}

	// value := 0
	// switch obj := result.(type) {
	// case model.Vector:
	// 	if obj.Len() == 0 {
	// 		return -1, errors.New("empty vector")
	// 	}
	// 	sample := obj[len(obj)-1]
	// 	value = int(sample.Value)
	// case *model.Scalar:
	// 	value = int(obj.Value)
	// }
	return result, nil
}

func (c *Client) generateClient() (api.Client, error) {
	roundTripper := api.DefaultRoundTripper
	if c.password != "" {
		roundTripper = config.NewBasicAuthRoundTripper(c.username, c.password, "", api.DefaultRoundTripper)
	}
	return api.NewClient(api.Config{
		Address:      c.url,
		RoundTripper: roundTripper,
	})
}
