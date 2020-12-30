package link_checker_events

import (
	"gitlab.com/vivsoft-platform/k8s-serverless/svc_kit/svc_pkg"
)

func Listen(url string, sink svc_pkg.LinkCheckerEvents) (err error) {
	conn, err := connect(url)
	if err != nil {
		return
	}

	conn.QueueSubscribe(subject, queue, func(e *Event) {
		sink.OnLinkChecked(e.Username, e.Url, e.Status)
	})

	return
}
