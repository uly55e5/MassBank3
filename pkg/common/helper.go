package common

import (
	"log"
	"os"
)

func GetEnv(name string, fallback string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}
	return fallback
}

func CheckAndLogError(err error, msg string) {
	if err != nil {
		log.Println(msg, ": ", err.Error())
	}
}
