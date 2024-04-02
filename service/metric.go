package service

import (
	"fmt"

	"github.com/DataDog/datadog-go/statsd"
)

const (
	retrieveFormat       = "%s.%s"
	descriptionTagFormat = "description:%s"
)

type MetricExporter interface {
	Count(name, description string, value int64, tags map[string]string) error
	Gauge(name, description string, value float64, tags map[string]string) error
	Histogram(name, description string, value float64, tags map[string]string) error
}

// MetricService is a service to collect metrics and send to datadog via statsd.Client.
type MetricService struct {
	client *statsd.Client
}

func NewMetricService(client *statsd.Client) MetricExporter {
	return &MetricService{
		client: client,
	}
}

// Count calls statsd.Client.Count() to send count on datadog.
func (t *MetricService) Count(name, description string, value int64, tags map[string]string) error {
	var tgs []string
	for key, tag := range tags {
		tgs = append(tgs, fmt.Sprintf(retrieveFormat, key, tag))
	}
	tgs = append(tgs, fmt.Sprintf(descriptionTagFormat, description))
	return t.client.Count(name, value, tgs, 1)
}

// Gauge calls statsd.Client.Gauge() to send gauge on datadog.
func (t *MetricService) Gauge(name, description string, value float64, tags map[string]string) error {
	var tgs []string
	for key, tag := range tags {
		tgs = append(tgs, fmt.Sprintf(retrieveFormat, key, tag))
	}
	tgs = append(tgs, fmt.Sprintf(descriptionTagFormat, description))
	return t.client.Gauge(name, value, tgs, 1)
}

// Histogram calls statsd.Client.Histogram() to send histogram on datadog.
func (t *MetricService) Histogram(name, description string, value float64, tags map[string]string) error {
	var tgs []string
	for key, tag := range tags {
		tgs = append(tgs, fmt.Sprintf(retrieveFormat, key, tag))
	}
	tgs = append(tgs, fmt.Sprintf(descriptionTagFormat, description))
	return t.client.Histogram(name, value, tgs, 1)
}
