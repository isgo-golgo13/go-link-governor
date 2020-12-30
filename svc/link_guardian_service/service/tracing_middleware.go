package service

import (
	"github.com/opentracing/opentracing-go"
	"gitlab.com/vivsoft-platform/k8s-serverless/svc_kit/svc_pkg"
)

// implement function to return ServiceMiddleware
func newTracingMiddleware(tracer opentracing.Tracer) linkManagerMiddleware {
	return func(next svc_pkg.LinkManager) svc_pkg.LinkManager {
		return tracingMiddleware{next, tracer}
	}
}

type tracingMiddleware struct {
	next   svc_pkg.LinkManager
	tracer opentracing.Tracer
}

func (m tracingMiddleware) GetLinks(request svc_pkg.GetLinksRequest) (result svc_pkg.GetLinksResult, err error) {
	defer func(span opentracing.Span) {
		span.Finish()
	}(m.tracer.StartSpan("GetLinks"))
	result, err = m.next.GetLinks(request)
	return
}

func (m tracingMiddleware) AddLink(request svc_pkg.AddLinkRequest) error {
	return m.next.AddLink(request)
}

func (m tracingMiddleware) UpdateLink(request svc_pkg.UpdateLinkRequest) error {
	return m.next.UpdateLink(request)
}

func (m tracingMiddleware) DeleteLink(username string, url string) error {
	return m.next.DeleteLink(username, url)
}
