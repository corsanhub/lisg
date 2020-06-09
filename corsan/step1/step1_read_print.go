package step1

import (
	"bufio"
	"fmt"
	"os"

	"corsanhub.com/lisg/corsan/core"
	"corsanhub.com/lisg/corsan/logging"
	"corsanhub.com/lisg/corsan/util"
)

var log = logging.Logger{Name: "core.reader"}

func READ(value string) core.MalType {
	readValue := core.XReadStr(value)
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

//Step1ReadPrint - Executes Step 1
func Step1ReadPrint() {
	for {
		rdr := bufio.NewReader(os.Stdin)
		fmt.Print("user> ")
		text, _ := rdr.ReadString('\n')
		textx := REP(text)
		fmt.Printf("%-v\n", textx)
	}
}
