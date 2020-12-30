package link_checker_events

import (
	"log"

	"github.com/nats-io/go-nats"
	"gitlab.com/vivsoft-platform/k8s-serverless/svc_kit/svc_pkg"
)

type eventSender struct {
	hostname string
	nats     *nats.EncodedConn
}

func (s *eventSender) OnLinkChecked(username string, url string, status svc_pkg.LinkStatus) {
	err := s.nats.Publish(subject, Event{username, url, status})
	if err != nil {
		log.Fatal(err)
	}
}

func NewEventSender(url string) (svc_pkg.LinkCheckerEvents, error) {
	ec, err := connect(url)
	if err != nil {
		return nil, err
	}
	return &eventSender{hostname: url, nats: ec}, nil
}
