package message

import (
	"errors"
	"sendmsg/global"
)

type Message interface {
	Send(body Body) error
}

type Body struct {
	Title   string
	Content string
}

func GetType() (Message, error) {
	switch global.GLO_CONF.MsgType {
	case "bark":
		if global.GLO_CONF.BarkUrl == "" || global.GLO_CONF.BarkKey == "" {
			return nil, errors.New("bark conf not completed")
		}
		return &Bark{
			url: global.GLO_CONF.BarkUrl,
			key: global.GLO_CONF.BarkKey,
		}, nil
	case "mail":
		if global.GLO_CONF.MailHost == "" || global.GLO_CONF.MailPort == 0 || global.GLO_CONF.MailUser == "" || global.GLO_CONF.MailPwd == "" || global.GLO_CONF.MailFromName == "" || global.GLO_CONF.MailTo == nil {
			return nil, errors.New("mail conf not completed")
		}
		return &Mail{
			Host:     global.GLO_CONF.MailHost,
			Port:     global.GLO_CONF.MailPort,
			Username: global.GLO_CONF.MailUser,
			Password: global.GLO_CONF.MailPwd,
			FromName: global.GLO_CONF.MailFromName,
			To:       global.GLO_CONF.MailTo,
		}, nil
	}
	return nil, errors.New("message type in config is not supported")
}

func Enabled() bool {
	return global.GLO_CONF.MsgEnabled
}
