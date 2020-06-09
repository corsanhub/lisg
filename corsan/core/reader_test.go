package core

import (
	"fmt"
	"runtime"
	"testing"

	"corsanhub.com/lisg/corsan/util"
)

func nameFromStack(level int) string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[level])
	return f.Name()
}

func getPassStr(input string, result string, expected string) string {
	testName := nameFromStack(2)
	return fmt.Sprintf("··· PASSED: %s\n|--💚Input    ->%s<-\n|--  Result   ->%s<-\n|--  Expected ->%s<-", testName, input, result, expected)
}

func getErrorStr(input string, result string, expected string) string {
	return fmt.Sprintf("Error tokenizing string.\n|--💔Input    ->%s<-\n|--  Result   ->%s<-\n|--  Expected ->%s<-", input, result, expected)
}

func probeTokenize(t *testing.T, input string, expected string) {
	tokens := XTokenize(input)
	result := util.PointersToString(tokens, "")

	if result != expected {
		errorStr := getErrorStr(input, result, expected)
		t.Errorf(errorStr)
	} else {
		succeedStr := getPassStr(input, result, expected)
		t.Logf(succeedStr)
	}
}

func TestTokenize(t *testing.T) {
	inputs := []string{
		"(let [y (some 3.4)\n  (print (+ 3 4))])\n(println \"Hello!\")",
		"(def x (inc 2))",
		"(inc 1)",
	}
	outputs := []string{
		"( let [ y ( some 3.4 ) ( print ( + 3 4 ) ) ] ) ( println \"Hello!\" ) ",
		"( def x ( inc 1 ) ) ",
		"( inc 1 ) ",
	}
	for i, input := range inputs {
		probeTokenize(t, input, outputs[i])
	}
}
