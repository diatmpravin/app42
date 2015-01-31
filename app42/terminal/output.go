package terminal

import (
	"fmt"
)

func Ask(prompt string) (ans string) {
	fmt.Printf(prompt + " ")
	fmt.Scanln(&ans)
	return
}

func Say(message string, args ...interface{}) {
	fmt.Printf(message+"\n", args...)
	return
}
