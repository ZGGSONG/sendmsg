package common

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var Log *log.Logger

// InitLog
//
//	@Description: 初始化原生 log
func InitLog() {
	path := "./log"
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = os.Mkdir(path, 0755)
	}
	//SavePath := fmt.Sprintf("./log/log_%v.log", time.Now().Format("20060102_150405"))
	SavePath := fmt.Sprintf("%v/log.log", path)
	printConsole := true
	for index, value := range os.Args {
		if value == "-log_path" {
			index += 1
			SavePath = os.Args[index]
		} else if value == "-log_ui" {
			index += 1
			va, err := strconv.Atoi(os.Args[index])
			if err != nil {
				fmt.Printf("cant parse -log_ui value ,current set value:", os.Args[index])
			} else {
				if va != 0 {
					printConsole = true
				} else {
					printConsole = false
				}
			}
		}
	}
	logFile, er := os.OpenFile(SavePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	Log = log.New(logFile, "", log.Ltime|log.Lshortfile|log.LstdFlags)
	if er != nil {
		fmt.Printf("log file init faild! can't open file:%s ,error msg:%s", SavePath, er)
	} else {
		if printConsole {
			mw := io.MultiWriter(os.Stdout, logFile)
			Log.SetOutput(mw)
		}
	}
}
