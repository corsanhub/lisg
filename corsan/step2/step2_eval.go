package step2

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"corsanhub.com/lisg/corsan/core"
	"corsanhub.com/lisg/corsan/fn"
	"corsanhub.com/lisg/corsan/logging"
	"corsanhub.com/lisg/corsan/util"
)

var log = logging.Logger{Name: "core.reader"}

func read(value string) core.MalType {
	readValue := core.ReadStr(value)
	log.Debug(util.Xs("readValue  : %v", readValue))
	return readValue
}

type multiFn func(args ...core.MalType) core.MalType

func eval(ast core.MalType, env map[string]interface{}) (interface{}, error) {
	log.Debug(util.Xs("AST value  : %v", ast))
	mType := core.GetType(ast)
	if mType != "core.MalList" {
		result, err := evalAst(ast, env)
		return result, err
	} else {
		childs := ast.Value().([]core.MalType)
		size := len(childs)
		if size == 0 {
			return ast, nil
		} else {
			newAst, _ := evalAst(ast, env)
			aa := newAst.(core.MalType)
			items := aa.Value().([]interface{})
			args := []core.MalType{}
			var fx = items[0]

			for _, item := range items[1:] {
				current := item.(core.MalType)
				args = append(args, current)
			}

			// copy(items[1:], args[:])

			result := fx.(multiFn)(args...)
			return result, nil
		}
	}
}

func evalAst(ast core.MalType, env map[string]interface{}) (interface{}, error) {
	mType := core.GetType(ast)

	if mType == "core.MalSymbol" {
		symbolName := ast.PrintStr()
		fx := env[symbolName]
		var err error = nil
		if fx == nil {
			f := util.TraceStr(0)
			err = core.NewError(f, fmt.Sprintf("undefined '%+v'", symbolName))
		}
		return fx, err
	} else if mType == "core.MalList" {
		list := ast.(core.MalList)
		childs := list.Value().([]core.MalType)
		for _, child := range childs {
			eval(child, env)
		}
	} else {
		return ast, nil
	}
	return ast, nil
}

func print(value core.MalType) string {
	printValue := core.PrintStr(value)
	log.Debug(util.Xs("printValue : %v", printValue))
	return printValue
}

func REP(str string, env map[string]interface{}) string {
	readResult := read(str)
	evalResult, _ := eval(readResult, env)
	prntResult := print(evalResult.(core.MalType))
	return prntResult
}

func InitializeCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\n~~~~> Exiting REPL ...")
		os.Exit(0)
	}()
}

var (
	historyFn = filepath.Join(os.TempDir(), "lish.repl")
	names     = []string{"john", "james", "mary", "nancy"}
)

//Step2Eval - Executes Step 2
func Step2Eval() {
	InitializeCloseHandler()

	var text = ""
	var err error

	env := map[string]interface{}{
		"+": fn.Sum,
	}

	for {
		rdr := bufio.NewReader(os.Stdin)
		fmt.Print("user> ")
		text, err = rdr.ReadString('\n')

		if err != nil {
			break
		}

		if text == "\n" {
			fmt.Println("")
		} else {
			resultText := REP(text, env)
			if resultText == "exit" {
				fmt.Println("\n~~~~> Exiting REPL ...")
				os.Exit(0)
			} else {
				fmt.Printf("%-v\n", resultText)
			}
		}
	}
}
