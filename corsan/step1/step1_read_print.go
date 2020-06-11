package step1

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"corsanhub.com/lisg/corsan/core"
	"corsanhub.com/lisg/corsan/logging"
	"corsanhub.com/lisg/corsan/util"
)

var log = logging.Logger{Name: "core.reader"}

func READ(value string) core.MalType {
	readValue := core.ReadStr(value)
	log.Debug(util.Xs("readValue  : %v", readValue))
	return readValue
}

func EVAL(value core.MalType) core.MalType {
	log.Debug(util.Xs("evalValue  : %v", value))
	return value
}

func PRINT(value core.MalType) string {
	printValue := core.PrintStr(value)
	log.Debug(util.Xs("printValue : %v", printValue))
	return printValue
}

func REP(str string) string {
	return PRINT(EVAL(READ(str)))
}

func InitializeCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\n-- Exiting REPL ...")
		os.Exit(0)
	}()
}

//Step1ReadPrint - Executes Step 1
func Step1ReadPrint() {
	InitializeCloseHandler()

	var text = ""
	var err error

	for {
		rdr := bufio.NewReader(os.Stdin)
		fmt.Print("user> ")
		text, err = rdr.ReadString('\n')

		if err != nil {
			break
		}

		resultText := REP(text)
		if resultText == "exit" {
			fmt.Println("\n-- Exiting REPL ...")
			os.Exit(0)
		} else {
			fmt.Printf("%-v\n", resultText)
		}
	}
}

func xx() {
	var text = ""
	var err error
	for {
		rdr := bufio.NewReader(os.Stdin)
		fmt.Print("user> ")
		text, err = rdr.ReadString('\n')
		// text, err = liner.Prompt("user> ")
		if err != nil {
			break
		}

		log.Debug(fmt.Sprintf("----- %#v", text))
		if text == "" {
			fmt.Println("")
		} else {
			resultText := REP(text)
			fmt.Printf("%-v\n", resultText)
		}
	}

}
