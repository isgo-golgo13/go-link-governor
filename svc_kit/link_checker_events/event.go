package link_checker_events

import (
	"gitlab.com/vivsoft-platform/k8s-serverless/svc_kit/svc_pkg"
)

type Event struct {
	Username string
	Url      string
	Status   svc_pkg.LinkStatus
}
