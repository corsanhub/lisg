package main

import (
	"bufio"
	"fmt"
	"os"
)

func add(x int, y int) int {
	return x + y
}

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

func repl() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("user> ")
		text, _ := reader.ReadString('\n')
		textx := rep(text)
		fmt.Printf("%s", textx)
	}
}

func main() {
	name := "Daj K'ptzin"
	fmt.Printf("Hello there %s!\n", name)
	repl()
}
