package step1

import (
	"bufio"
	"fmt"
	"os"

	"corsanhub.com/lisg/corsan/core"
)

func READ(value string) core.MalType {
	return nil
}

func EVAL(value core.MalType) core.MalType {
	return nil
}

func PRINT(value core.MalType) string {
	return "value"
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
