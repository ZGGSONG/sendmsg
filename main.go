package main

import (
	log "github.com/sirupsen/logrus"
	"sendmsg/common"
	"sendmsg/global"
	"sendmsg/message"
	"sendmsg/service"
	"time"
)

func init() {
	/*日志初始化*/
	common.InitLog2()

	/*配置初始化*/
	conf, err := common.InitConfig()
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}
	global.GLO_CONF = conf

}
func main() {
	/*测试发送*/
	var s = service.Service{Body: message.Body{
		Title:   "test title",
		Content: "this is content, time is " + time.Now().Format("2006-01-02 15:04:05"),
	}}
	if err := s.Run(); err != nil {
		log.Fatalf("failed to send message: %v", err)
	} else {
		log.Printf("send message successfully...")
	}

	/*监听配置*/
	for {
		_conf := <-global.GLO_CONF_CH
		log.Printf("config changed: %v", _conf)
		global.GLO_CONF = _conf
	}
}
