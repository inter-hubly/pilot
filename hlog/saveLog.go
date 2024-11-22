package hlog

import (
	"log"
	"os"
)

func SaveLog(filename string, message string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	logger.Println(message)

	return nil
}
