package service

import (
	"time"

	"github.com/go-kit/kit/log"
	"gitlab.com/vivsoft-platform/k8s-serverless/svc_kit/svc_pkg"
)

// implement function to return ServiceMiddleware
func newLoggingMiddleware(logger log.Logger) linkManagerMiddleware {
	return func(next svc_pkg.LinkManager) svc_pkg.LinkManager {
		return loggingMiddleware{next, logger}
	}
}

type loggingMiddleware struct {
	next   svc_pkg.LinkManager
	logger log.Logger
}

func (m loggingMiddleware) GetLinks(request svc_pkg.GetLinksRequest) (result svc_pkg.GetLinksResult, err error) {
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "GetLinks",
			"request", request,
			"result", result,
			"duration", time.Since(begin),
		)
	}(time.Now())
	result, err = m.next.GetLinks(request)
	return
}

func (m loggingMiddleware) AddLink(request svc_pkg.AddLinkRequest) error {
	return m.next.AddLink(request)
}

func (m loggingMiddleware) UpdateLink(request svc_pkg.UpdateLinkRequest) error {
	return m.next.UpdateLink(request)
}

func (m loggingMiddleware) DeleteLink(username string, url string) error {
	return m.next.DeleteLink(username, url)
}
