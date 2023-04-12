package helper

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitEnv(filename string) {

	homedir, _ := os.UserHomeDir()
	envdir := homedir + "/.gas/envs/"
	f := envdir + "/" + filename

	if _, err := os.Stat(f); err == nil {
		err := godotenv.Load(f) //read .evn file
		if err != nil {
			log.Fatal("Error loading .env file ", f)
		}
	} else {
		err := os.MkdirAll(envdir, os.ModePerm) //create dir
		if err != nil {
			log.Println(err)
		}

		err = os.WriteFile(f, []byte(""), 0644)
		if err != nil {
			log.Println("Cannot create ", f)
		}

	}

}
