package main

import (
	"log"
	"sendmsg/common"
	"sendmsg/global"
	"sendmsg/message"
	"sendmsg/model"
	"sendmsg/service"
	"time"
)

func init() {
	/*日志初始化*/
	common.InitLog()

	/*配置初始化*/
	global.GLO_CONF_CH = make(chan model.Config)
	_conf, err := common.InitConfig()
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
			if count > 1 {
				return
			}
			count++
			var s = service.Service{Body: message.Body{
				Title:   "test title",
				Content: "this is content, time is " + time.Now().Format("2006-01-02 15:04:05"),
			}}
			s.Run()
			time.Sleep(time.Second * 3)
		}
	}()

	for {
		_conf := <-global.GLO_CONF_CH
		global.GLO_CONF = _conf
	}
}
