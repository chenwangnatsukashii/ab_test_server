package utils

import (
	"log"
	"os"
	"time"
)

func LogInit() {
	// 获取日志文件句柄，以只写入文件|没有时创建|文件尾部追加 的形式打开这个文件
	logFile, err := os.OpenFile("./log/system/"+time.Now().Format("2006-01-02")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	// 设置存储位置
	log.SetOutput(logFile)
}
