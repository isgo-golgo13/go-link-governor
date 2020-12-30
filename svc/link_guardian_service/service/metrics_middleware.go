package service

import (
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"gitlab.com/vivsoft-platform/k8s-serverless/svc_kit/metrics"
	"gitlab.com/vivsoft-platform/k8s-serverless/svc_kit/svc_pkg"
)

// implement function to return ServiceMiddleware
func newMetricsMiddleware() linkManagerMiddleware {
	return func(next svc_pkg.LinkManager) svc_pkg.LinkManager {
		m := metricsMiddleware{next,
			map[string]prometheus.Counter{},
			map[string]prometheus.Summary{}}
		methodNames := []string{"GetLinks", "AddLink", "UpdateLink", "DeleteLink"}
		for _, name := range methodNames {
			m.requestCounter[name] = metrics.NewCounter("link", strings.ToLower(name)+"_count", "count # of requests")
			m.requestLatency[name] = metrics.NewSummary("link", strings.ToLower(name)+"_summary", "request summary in milliseconds")

		}
		return m
	}
}

type metricsMiddleware struct {
	next           svc_pkg.LinkManager
	requestCounter map[string]prometheus.Counter
	requestLatency map[string]prometheus.Summary
}

func (m metricsMiddleware) recordMetrics(name string, begin time.Time) {
	m.requestCounter[name].Inc()
	durationMilliseconds := float64(time.Since(begin).Nanoseconds() * 1000000)
	m.requestLatency[name].Observe(durationMilliseconds)
}

func (m metricsMiddleware) GetLinks(request svc_pkg.GetLinksRequest) (result svc_pkg.GetLinksResult, err error) {
	defer func(begin time.Time) {
		m.recordMetrics("GetLinks", begin)
	}(time.Now())
	result, err = m.next.GetLinks(request)
	return
}

func (m metricsMiddleware) AddLink(request svc_pkg.AddLinkRequest) error {
	defer func(begin time.Time) {
		m.recordMetrics("AddLink", begin)
	}(time.Now())
	return m.next.AddLink(request)
}

func (m metricsMiddleware) UpdateLink(request svc_pkg.UpdateLinkRequest) error {
	defer func(begin time.Time) {
		m.recordMetrics("UpdateLink", begin)
	}(time.Now())
	return m.next.UpdateLink(request)
}

func (m metricsMiddleware) DeleteLink(username string, url string) error {
	defer func(begin time.Time) {
		m.recordMetrics("DeleteLink", begin)
	}(time.Now())
	return m.next.DeleteLink(username, url)
}
