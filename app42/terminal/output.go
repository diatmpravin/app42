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

func Failed(message string, err error) {
	Say(Red("FAILED"))

	if message != "" {
		Say(message)
	}

	if err != nil {
		Say(err.Error())
	}
	return
}

func Ok() {
	Say(Green("OK"))
}
