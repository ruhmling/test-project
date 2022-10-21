package log

import "fmt"

func Debug(message string, args ...string) {
	println(fmt.Sprintf(message, args))
}
