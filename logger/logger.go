package logger

import (
	"fmt"
	"log"
)

var colorReset = string("\033[0m")

func Log(id int32, message string, a ...any) {
	log.Println(
		fmt.Sprintf("\033[%dm", id), // colorize
		fmt.Sprintf(message, a...),  // message
		colorReset,                  // color reset
	)
}
