package core

import (
	"fmt"
	"regexp"
)

var tokenRegexStr = "[\\s,]*(~@|[\\[\\]{}()'`~^@]|\"(?:\\\\.|[^\\\\\"])*\"?|;.*|[^\\s\\[\\]{}('\"`,;)]*)"
var intRegexStr = `^[0-9]+$`
var floatRegexStr = `^\d*[.]\d+$`

var stringRegex = regexp.MustCompile(tokenRegexStr)
var intRegex = regexp.MustCompile(intRegexStr)
var floatRegex = regexp.MustCompile(floatRegexStr)

func Tokenize(str string) []string {
	tokens := make([]string, 0, 1)
	for _, group := range stringRegex.FindAllStringSubmatch(str, -1) {
		if (group[1] == "") || (group[1][0] == ';') {
			continue
		}
		tokens = append(tokens, group[1])
	}
	return tokens
}

// func printFnStart(fn string, reader **Reader, value string) {
// 	// ss := fmt.Sprintf("[%+v] %+s              => %+s <=", reader, fn, value)
// 	// fmt.Println(ss)
// }

// type Reader struct {
// 	position int
// 	tokens   []string
// }

// func (reader *Reader) next() string {
// 	currentToken := reader.tokens[reader.position]
// 	reader.position++
// 	return currentToken
// }

// func (reader *Reader) peek() string {
// 	currentToken := reader.tokens[reader.position]
// 	return currentToken
// }

// func read_atom(reader *Reader) MalType {
// 	token := reader.next()
// 	// printFnStart("read_atom", &reader, reader.token)

// 	intMatches := intRegex.Match([]byte(*token))
// 	floatMatches := floatRegex.Match([]byte(*token))

// 	if intMatches {
// 		fmt.Println("intMatches:", intMatches)
// 		intValue, _ := strconv.ParseInt(*token, 10, 64)
// 		return MalInteger{value: intValue}
// 	} else if floatMatches {
// 		fmt.Println("intMatches:", intMatches)
// 		floatValue, _ := strconv.ParseFloat(*token, 10)
// 		return &MalFloat{value: floatValue}
// 	} else {
// 		if token != nil {
// 			return &MalObject{value: token}
// 		} else {
// 			return nil
// 		}
// 	}
// }

// func read_list(reader *Reader) MalType {
// 	token := reader.next()

// 	printFnStart("read_list", &reader, *token)

// 	elements := []*MalType{}
// 	for {
// 		current := reader.next()
// 		var obj = &MalObject{}
// 		if current.token[0] == ')' {
// 			break
// 		} else {
// 			obj = read_atom(current)
// 		}

// 		elements = append(elements, obj)
// 	}

// 	list := &MalList{
// 		elements: elements,
// 	}
// 	return &list.MalObject
// }

// func read_form(reader *Reader) *MalObject {
// 	token := reader.peek()
// 	printFnStart("read_form", &reader, reader.token)

// 	//fmt.Println(fmt.Sprintf("current reader: [%+v]", &current))

// 	var obj = &MalObject{}
// 	switch token[0] {
// 	case '(':
// 		obj = reader.read_list()
// 	default:
// 		obj = reader.read_atom()
// 	}
// 	return obj
// }

// // func ppstr(items []string) string {
// // 	var buffer bytes.Buffer
// // 	for _, item := range items {
// // 		buffer.WriteString("°" + item + "° ")
// // 	}
// // 	return buffer.String()
// // }

// func Read_str(tokenStr string) *Reader {
// 	reader := &Reader{}
// 	//printFnStart("read_str ", &reader, tokenStr)
// 	tokens := tokenize(tokenStr)
// 	reader.tokens = tokens
// 	fmt.Println("read_str:", reader.token)
// 	reader.position = 0

// 	readers := []*Reader{}
// 	if len(tokens) > 1 {
// 		for _, token := range tokens {
// 			xtoken := strings.TrimSpace(token)
// 			newReader := Read_str(xtoken)

// 			readers = append(readers, newReader)
// 		}
// 		reader.readers = readers

// 		reader.read_form()
// 	}

// 	fmt.Println("reader:", reader.token)
// 	return reader
// }

func TestReader() {
	//fmt.Printf("Position: %d, Token: %+v\n", currentPosition, currentToken)
	println("Testing reader ...")
	println("---------------------------------------------------------------")

	//fmt.Println("regexpStr:", regexpStr)

	str := "(let [y (some 3.4)\n  (print (+ 3 4))])\n(println \"Hello!\")"
	str = "(def x (inc 1))"
	fmt.Println("str: ", str)

	// reader := Read_str(str)

	// fmt.Println(": MAIN READER :", reader)

	// for _, token := range tokens {
	// 	fmt.Println("token: ", token)
	// }

	println("---------------------------------------------------------------")

}
