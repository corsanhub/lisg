package step0

import (
	"bufio"
	"fmt"
	"os"
)

func READ(value string) string {
	return value
}

func EVAL(value string) string {
	return value
}

func PRINT(value string) string {
	return value
}

func rep(text string) string {
	readResult := READ(text)
	evalResult := EVAL(readResult)
	printResult := PRINT(evalResult)

	return printResult
}

//Step0Repl - Executes Step 0
func Step0Repl() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("user> ")
		text, _ := reader.ReadString('\n')
		textx := rep(text)
		fmt.Printf("%s", textx)
	}
}
