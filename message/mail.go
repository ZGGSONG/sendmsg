package message

import "log"

type Mail struct {
	Host     string
	Protocol string
	Port     int
	Username string
	Password string
	FromName string
	To       []string
}

func (m *Mail) Send(body Body) {
	log.Printf("[mail] sending message...")
	log.Printf(body.Title + " " + body.Content)
}
