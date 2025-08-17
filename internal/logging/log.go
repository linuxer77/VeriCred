package logging

import (
	"log"
	"os"
)
var Logger *log.Logger

func Init()  {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	Logger = log.New(file, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Println("Logger initialized")
}