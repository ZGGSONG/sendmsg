package service

import (
	"log"
	"sendmsg/message"
)

type Support struct{}

func (s Support) Run() {
	log.Printf("[support] start excute...")

}

func send(body message.Body) {
	m := message.GetType()
	if m == nil || !message.Enabled() {
		return
	}
	m.Send(body)
}
