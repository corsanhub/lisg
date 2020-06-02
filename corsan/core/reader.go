package core

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var tokenRegexStr = "[\\s,]*(~@|[\\[\\]{}()'`~^@]|\"(?:\\\\.|[^\\\\\"])*\"?|;.*|[^\\s\\[\\]{}('\"`,;)]*)"
var intRegexStr = `^[0-9]+$`

var stringRegex = regexp.MustCompile(tokenRegexStr)
var intRegex = regexp.MustCompile(intRegexStr)

type Reader struct {
	position int
	token    string
	readers  []*Reader
}

func printFnStart(fn string, reader **Reader, value string) {
	ss := fmt.Sprintf("[%+v] %+s              => %+s <=", reader, fn, value)
	fmt.Println(ss)
}

func (reader *Reader) next() *Reader {
	n_reader := reader.readers[reader.position]
	reader.position++
	return n_reader
}

func (reader *Reader) peek() *Reader {
	n_reader := reader.readers[reader.position]
	return n_reader
}

func (reader *Reader) tokenize(str string) []string {

	matches := stringRegex.FindAllString(str, -1)
	return matches
}

func (reader *Reader) read_atom() *MalObject {
	printFnStart("read_atom", &reader, reader.token)
	//obj := &MalObject{str: reader.token}

	return nil
}

func (reader *Reader) read_list() *MalObject {
	printFnStart("read_list", &reader, reader.token)

	elements := []*MalObject{}
	for {
		current := reader.next()
		var obj = &MalObject{}
		if current.token[0] == ')' {
			break
		} else {
			obj = current.read_atom()
		}

		elements = append(elements, obj)
	}

	list := &MalList{
		elements: elements,
	}
	return &list.MalObject
}

func (reader *Reader) read_form() *MalObject {
	printFnStart("read_form", &reader, reader.token)

	current := reader.peek()
	//fmt.Println(fmt.Sprintf("current reader: [%+v]", &current))
	token := current.token

	var obj = &MalObject{}
	switch token[0] {
	case '(':
		obj = reader.read_list()
	default:
		obj = reader.read_atom()
	}
	return obj
}

func ppstr(items []string) string {
	var buffer bytes.Buffer
	for _, item := range items {
		buffer.WriteString("°" + item + "° ")
	}
	return buffer.String()
}

func read_str(tokenStr string) *Reader {
	reader := &Reader{
		token: tokenStr,
	}

	printFnStart("read_str ", &reader, tokenStr)

	tokens := reader.tokenize(tokenStr)
	reader.position = 0

	readers := []*Reader{}
	if len(tokens) > 1 {
		for _, token := range tokens {
			xtoken := strings.TrimSpace(token)
			newReader := read_str(xtoken)

			readers = append(readers, newReader)
		}
		reader.readers = readers

		reader.read_form()
	}

	fmt.Println("reader:", reader)
	return reader
}

func DoSomething() {
	//fmt.Printf("Position: %d, Token: %+v\n", currentPosition, currentToken)
	println("Doing something ...")
	println("---------------------------------------------------------------")

	//fmt.Println("regexpStr:", regexpStr)

	str := "(let [y (some 3.4)\n  (print (+ 3 4))])\n(println \"Hello!\")"
	str = "(def x (inc 1))"
	reader := read_str(str)

	fmt.Println(": MAIN READER :", reader)

	// for _, token := range tokens {
	// 	fmt.Println("token: ", token)
	// }

	println("---------------------------------------------------------------")

	//TestTypes()

	println("---------------------------------------------------------------")

}
