package core

import (
	"fmt"

	"corsanhub.com/lisg/corsan/logging"
	"corsanhub.com/lisg/corsan/util"
)

var log = logging.Logger{Name: "core.reader"}

func NewError(fnName, str string) *MalError {
	return &MalError{f: fnName, e: str}
}

type Reader struct {
	position int
	counter  int
	tokens   []*string
}

func CreateReader(str string) *Reader {
	tokens := Tokenize(str)
	log.Debug(util.Xs("tokens: %s\n", util.PointersToString(tokens, "'")))
	reader := &Reader{tokens: tokens}
	return reader
}

func ReadStr(str string) MalType {
	if &str != nil {
		reader := CreateReader(str)
		form, _ := reader.readForm()

		if reader.counter != 0 {
			return MalObject{v: "unbalanced"}
		}
		return form
	} else {
		return nil
	}
}

func TestXReader() {
	println("Testing reader ...")
	println("------------------------------------------------------------------------------------------------------------------------------")

	str := ""
	// str = "(a b)"
	// str = "(()())"
	// str = "^{\"a\" 1} [1 2 3]"
	// str = "'(1 2 3)"
	// str = "~(1 2 3)"
	// str = "@(a {:b 1})"
	str = "`(1 ~a 3)"
	reader := CreateReader(str)
	form, _ := reader.readForm()
	fmt.Printf("form: %#v\n", form)

	println("------------------------------------------------------------------------------------------------------------------------------")

}
