package core

import (
	"fmt"
	"regexp"
	"strconv"

	"corsanhub.com/lisg/corsan/util"
)

var tokenRegexStr = "[\\s,]*(~@|[\\[\\]{}()'`~^@]|\"(?:\\\\.|[^\\\\\"])*\"?|;.*|[^\\s\\[\\]{}('\"`,;)]*)"
var intRegexStr = `^[0-9]+$`
var floatRegexStr = `^\d*[.]\d+$`

var StringRegex = regexp.MustCompile(tokenRegexStr)
var IntRegex = regexp.MustCompile(intRegexStr)
var FloatRegex = regexp.MustCompile(floatRegexStr)

func Tokenize(str string) []string {
	matches := StringRegex.FindAllStringSubmatch(str, -1)
	tokens := make([]string, 0)
	for _, g := range matches {
		tokens = append(tokens, g[1])
	}
	return tokens
}

// func printFnStart(fn string, reader **Reader, value string) {
// 	// ss := fmt.Sprintf("[%+v] %+s              => %+s <=", reader, fn, value)
// 	// fmt.Println(ss)
// }

func SomeError(fnName, str string) *MalError {
	return &MalError{f: fnName, e: str}
}

type Reader struct {
	position int
	tokens   []string
}

const NILS = ""

func (reader *Reader) next() (string, error) {
	//fmt.Printf("position: %d, len: %d\n\n", reader.position, len(reader.tokens))
	if reader.position < len(reader.tokens) {
		currentToken := reader.tokens[reader.position]
		reader.position++
		return currentToken, nil
	} else {
		fn := util.GetFrame(0).Function
		errStr := "Changos"
		err := SomeError(fn, errStr)
		return NILS, err
	}

}

func (reader *Reader) peek() string {
	currentToken := reader.tokens[reader.position]
	return currentToken
}

func (reader *Reader) readAtom() MalType {
	token, _ := reader.next()
	//fmt.Println("atom token:", token)

	// printFnStart("read_atom", &reader, reader.token)

	intMatches := IntRegex.Match([]byte(token))
	floatMatches := FloatRegex.Match([]byte(token))

	if intMatches {
		//fmt.Println("intMatches:", intMatches)
		intValue, _ := strconv.ParseInt(token, 10, 64)
		return MalInteger{v: intValue}
	} else if floatMatches {
		//fmt.Println("floatMatches:", floatMatches)
		floatValue, _ := strconv.ParseFloat(token, 10)
		return MalFloat{v: floatValue}
	} else {
		if &token != nil {
			return MalSymbol{v: token}
		} else {
			return nil
		}
	}
}

func (reader *Reader) readList() MalType {
	//printFnStart("read_list", &reader, *token)

	list := MalList{}
	for {
		current, err := reader.next()
		//fmt.Println("list token:", current)
		if err != nil {
			//fmt.Println("An error has ocurred --> ", err.Error())
			break
		}

		letter := current[0]
		switch letter {
		case ')':
			break
		default:
			//fmt.Println("current:", current)

			atom := reader.readAtom()
			//fmt.Println(fmt.Printf("Apendding atom to list: %+v\n", atom))
			list.v = append(list.v, atom)
		}
	}

	return list
}

func (reader *Reader) readForm() MalType {
	token := reader.peek()
	//fmt.Println("form token:", token)
	//printFnStart("read_form", &reader, reader.token)

	//fmt.Println(fmt.Sprintf("current reader: [%+v]", &current))

	switch token[0] {
	case '(':
		list := reader.readList()
		return list
	default:
		atom := reader.readAtom()
		fmt.Println("atom:", atom)
		return atom
	}
}

func CreateReader(tokenStr string) *Reader {
	reader := &Reader{}
	//printFnStart("read_str ", &reader, tokenStr)
	tokens := Tokenize(tokenStr)
	//fmt.Println("===== tokens:", tokens)
	//fmt.Println("===== tokenx:", util.ArrayToString(tokens, "\""))
	reader.tokens = tokens
	reader.position = 0
	return reader
}

func ReadStr(str string) MalType {
	reader := CreateReader(str)
	form := reader.readForm()
	return form
}

func TestReader() {
	//fmt.Printf("Position: %d, Token: %+v\n", currentPosition, currentToken)
	println("Testing reader ...")
	println("---------------------------------------------------------------")

	//fmt.Println("regexpStr:", regexpStr)

	str := "(let [y (some 3.4)\n  (print (+ 3 4))])\n(println \"Hello!\")"
	str = "(def x (inc 1))"
	str = "(inc 1)"
	str = "(4)"
	fmt.Println("str: ", str)

	reader := CreateReader(str)
	form := reader.readForm()

	fmt.Println("reader :", reader)
	fmt.Println("form   :", form)

	// for _, token := range tokens {
	// 	fmt.Println("token: ", token)
	// }
	println("---------------------------------------------------------------")

}
