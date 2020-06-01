package main

import (
	"fmt"
	"os"
	"strings"

	"corsanhub.com/lisg/corsan/core"
	"corsanhub.com/lisg/corsan/step0"
	"corsanhub.com/lisg/corsan/step1"
)

func main() {
	//argsWithProg := os.Args
	//argsWithoutProg := os.Args[1:]
	//runnable := os.Args[0]

	if len(os.Args) != 2 {
		panic("No arguments has been supplied.")
	}

	step := os.Args[1]
	fmt.Printf("step: %s!\n", step)

	name := "Daj K'ptzin"
	fmt.Printf("Hello there %s!\n", name)
	fmt.Printf("I'm glad to inform you, we are now executing testings for [%s]\n", step)

	if strings.HasPrefix(step, "step0") {
		step0.Step0Repl()
	} else if strings.HasPrefix(step, "step1") {
		step1.Step1ReadPrint()
	} else {
		core.DoSomething()
	}

	println("That's all FOLKS!!")
}
