package step1

import (
	"bufio"
	"fmt"
	"os"

	"corsanhub.com/lisg/corsan/core"
)

func READ(value string) core.MalType {
	return core.ReadStr(value)
}

func EVAL(value core.MalType) core.MalType {
	return value
}

func PRINT(value core.MalType) core.MalType {
	return value
}

func rep(text string) core.MalType {
	readResult := READ(text)
	evalResult := EVAL(readResult)
	printResult := PRINT(evalResult)

	return printResult
}

//Step1ReadPrint - Executes Step 1
func Step1ReadPrint() {
	for {
		rdr := bufio.NewReader(os.Stdin)
		fmt.Print("user> ")
		text, _ := rdr.ReadString('\n')
		textx := rep(text)
		fmt.Printf("%-v\n", textx)
	}
}