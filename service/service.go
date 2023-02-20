package service

import (
	"errors"
	"sendmsg/message"
)

type Service struct {
	Body message.Body
}

func (s Service) Run() error {
	return send(s.Body)
}

func send(body message.Body) error {
	m, err := message.GetType()
	if err != nil {
		return err
	}
	if !message.Enabled() {
		return errors.New("config set message not enabled")
	}
	if err := m.Send(body); err != nil {
		return err
	}
	return nil
}
