package console

import (
	"fmt"
)

type Color string

const (
	colorRed          = "\u001b[31m"
	colorGreen        = "\u001b[32m"
	colorYellow       = "\u001b[33m"
	colorBlue         = "\u001b[34m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message)
}

func Info(message string) {
	colorize(colorBlue, message)
}

func Success(message string) {
	colorize(colorGreen, message)
}

func Warning(message string) {
	colorize(colorYellow, message)
}

func Error(message string) {
	colorize(colorRed, message)
}
