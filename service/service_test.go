package service

import (
	"sendmsg/config"
	"sendmsg/global"
	"sendmsg/message"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	conf, err := config.InitConfig()
	if err != nil {
		t.Fatalf("Failed to initialize config: %v", err)
	}
	global.GLO_CONF = conf
	var s = Service{Body: message.Body{
		Title:   "test title",
		Content: "this is content, time is " + time.Now().Format("2006-01-02 15:04:05"),
	}}
	s.Run()
}
