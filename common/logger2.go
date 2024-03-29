package common

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLog2
//
//	@Description: 初始化 logrus
func InitLog2() {
	logger := &lumberjack.Logger{
		Filename: "./log/log.log",
		//Filename:   "./log/log" + time.Now().Format("20060102_150405") + ".log",
		MaxSize:    100,  // 日志文件大小，单位是 MB
		MaxBackups: 10,   // 最大过期日志保留个数
		MaxAge:     28,   // 保留过期文件最大时间，单位 天
		Compress:   true, // 是否压缩日志，默认是不压缩。这里设置为true，压缩日志
		LocalTime:  true, // 是否使用本地时间，默认是使用UTC时间
	}

	log.SetOutput(logger) // logrus 设置日志的输出方式
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
