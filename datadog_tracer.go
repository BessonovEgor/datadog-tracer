package datadogtracer

import (
	"fmt"

	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const retrieveFormat = "%s.%s"

type Tracer interface {
	TraceData(request string) func()
}

type DataDogTracer struct {
	serviceName string
}

// NewTracer creates a new DataDogTracer instance.
func NewTracer(serviceName string) Tracer {
	return &DataDogTracer{serviceName: serviceName}
}

// TraceData default trace function.
func (t *DataDogTracer) TraceData(sql string) func() {
	span := ddTracer.StartSpan(fmt.Sprintf(retrieveFormat, t.serviceName, sql))
	spanFinish := func() {
		span.Finish()
	}
	return spanFinish
}
