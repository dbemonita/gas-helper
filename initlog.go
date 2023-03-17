package helper

import (
	"log"
	"os"
)

func InitLog(filename string) {
	homedir, _ := os.UserHomeDir()

	logdir := homedir + "/.gas/logs/"

	err := os.MkdirAll(logdir, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	logFile, err := os.OpenFile(logdir+"/"+filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
