package service

import (
	"sendmsg/common"
	"sendmsg/message"
)

type Service struct {
	Body message.Body
}

func (s Service) Run() {
	common.Log.Printf("[service] start excute...")
	send(s.Body)
}

func send(body message.Body) {
	m := message.GetType()
	if m == nil || !message.Enabled() {
		return
	}
	m.Send(body)
}
