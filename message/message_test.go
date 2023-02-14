package message

import (
	"log"
	"sendmsg/config"
	"sendmsg/global"
	"testing"
)

func TestSend(t *testing.T) {
	conf, err := config.InitConfig()
	if err != nil {
		log.Println("[test] Failed to initialize config: ", err)
	}
	global.GLO_CONF = conf

	m := GetType()
	if !Enabled() {
		return
	}
	m.Send(Body{
		Title:   "Hello",
		Content: "World",
	})
}
