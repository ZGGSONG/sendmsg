package service

import (
	log "github.com/sirupsen/logrus"
	"sendmsg/message"
)

type Service struct {
	Body message.Body
}

func (s Service) Run() {
	log.Printf("[service] start excute...")
	send(s.Body)
}

func send(body message.Body) {
	m := message.GetType()
	if m == nil || !message.Enabled() {
		return
	}
	m.Send(body)
}
