package util

import (
	"log"
	"os"
)

func AppendLineToFile(logFile string, str string) {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(str + "\n"); err != nil {
		log.Println(err)
	}
}
