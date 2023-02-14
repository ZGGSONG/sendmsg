package main

import (
	"log"
	"sendmsg/config"
	"sendmsg/global"
	"sendmsg/message"
	"sendmsg/model"
	"sendmsg/service"
	"time"
)

func init() {
	global.GLO_CONF_CH = make(chan model.Config)
	_conf, err := config.InitConfig()
	if err != nil {
		log.Fatalf("[init] Failed to initialize config: %v", err)
	}
	global.GLO_CONF = _conf
}
func main() {
	//test
	go func() {
		var count int
		for {
			if count > 5 {
				return
			}
			count++
			var s = service.Service{Body: message.Body{
				Title:   "test title",
				Content: "this is content, time is " + time.Now().Format("2006-01-02 15:04:05"),
			}}
			s.Run()
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		_conf := <-global.GLO_CONF_CH
		global.GLO_CONF = _conf
	}
}
