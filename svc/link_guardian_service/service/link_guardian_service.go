package service

import (
	"fmt"
	"io"

	_ "github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/uber/jaeger-client-go"
	jeagerconfig "github.com/uber/jaeger-client-go/config"

	"gitlab.com/vivsoft-platform/k8s-serverless/svc_kit/svc_pkg"
)

type EventSink struct {
}

type linkManagerMiddleware func(svc_pkg.LinkManager) svc_pkg.LinkManager

func (s *EventSink) OnLinkAdded(username string, link *svc_pkg.Link) {
	//log.Println("Link added")
}

func (s *EventSink) OnLinkUpdated(username string, link *svc_pkg.Link) {
	//log.Println("Link updated")
}

func (s *EventSink) OnLinkDeleted(username string, url string) {
	//log.Println("Link deleted")
}

// createTracer returns an instance of Jaeger Tracer that samples
// 100% of traces and logs all spans to stdout.
func createTracer(service string) (opentracing.Tracer, io.Closer) {
	cfg := &jeagerconfig.Configuration{
		ServiceName: service,
		Sampler: &jeagerconfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jeagerconfig.ReporterConfig{
			LogSpans: true,
		},
	}
	logger := jeagerconfig.Logger(jaeger.StdLogger)
	tracer, closer, err := cfg.NewTracer(logger)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot create tracer: %v\n", err))
	}
	return tracer, closer
}

func Run() {

	// Create a logger
	logger := log.NewLogger("link manager")

	// Create a tracer
	tracer, closer := createTracer("link-manager")
	defer closer.Close()

	// Create the service implementation

}
