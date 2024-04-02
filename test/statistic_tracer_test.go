package test

import (
	"log"
	"testing"

	"github.com/BessonovEgor/datadog-tracer/service"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/stretchr/testify/require"
)

const (
	ddAddress = `localhost:8125`
)

func TestCount(t *testing.T) {
	client, err := statsd.New(ddAddress)
	defer func(client *statsd.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(client)
	require.NoError(t, err)

	client.Tags = []string{"environment:test"}

	srv := service.NewMetricService(client)
	err = srv.Count("test_count", "A test count metric from open telemetry collector.", 2, map[string]string{"kind": "test"})
	require.NoError(t, err)
}

func TestGauge(t *testing.T) {
	client, err := statsd.New(ddAddress)
	defer func(client *statsd.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(client)
	require.NoError(t, err)

	client.Tags = []string{"environment:test"}

	srv := service.NewMetricService(client)
	err = srv.Gauge("test_gauge", "A test gauge metric from open telemetry collector.", 5, map[string]string{"kind": "test"})
	require.NoError(t, err)
}

func TestHistogram(t *testing.T) {
	client, err := statsd.New(ddAddress)
	defer func(client *statsd.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(client)
	require.NoError(t, err)

	client.Tags = []string{"environment:test"}

	srv := service.NewMetricService(client)
	err = srv.Histogram("test_histogram", "A test histogram metric from open telemetry collector.", 2, map[string]string{"kind": "test"})
	require.NoError(t, err)
}
