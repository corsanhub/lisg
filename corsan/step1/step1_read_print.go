package step1

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

//Step1ReadPrint - Executes Step 1
func Step1ReadPrint() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("user> ")
		text, _ := reader.ReadString('\n')
		textx := rep(text)
		fmt.Printf("%s", textx)
	}
}
