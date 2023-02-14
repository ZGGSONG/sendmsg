package message

import (
	"log"
	"sendmsg/config"
	"sendmsg/global"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	conf, err := config.InitConfig()
	if err != nil {
		log.Println("[test] Failed to initialize config: ", err)
	}
	global.GLO_CONF = conf

	m := GetType()
	if !Enabled() || m == nil {
		t.Fatalf("[test] Failed to get type...")
	}
	m.Send(Body{
		Title:   "Hello",
		Content: "World " + time.Now().Format("2006-01-02 15:04:05"),
	})
}
