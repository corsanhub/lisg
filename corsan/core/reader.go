package core

import (
	"fmt"
	"regexp"
	"strconv"

	"corsanhub.com/lisg/corsan/logging"
	"corsanhub.com/lisg/corsan/util"
)

var log = logging.Logger{Name: "core.reader"}

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

func SomeError(fnName, str string) *MalError {
	return &MalError{f: fnName, e: str}
}

type Reader struct {
	position int
	counter  int
	tokens   []string
}

const NILS = ""

func (reader *Reader) next() (string, error) {
	if reader.position < len(reader.tokens) {
		currentToken := reader.tokens[reader.position]
		log.Debug(util.Xs("reader.position: %v", reader.position))
		log.Debug(util.Xs("currentToken   : %v", currentToken))
		reader.position++
		return currentToken, nil
	} else {
		fn := util.TraceStr(0)
		errStr := util.Xs("Out of bounds error %d", reader.position)
		err := SomeError(fn, errStr)
		return NILS, err
	}
}

func (reader *Reader) peek() (string, error) {
	if reader.position < len(reader.tokens) {
		currentToken := reader.tokens[reader.position]
		log.Debug(util.Xs("reader.position: %v", reader.position))
		log.Debug(util.Xs("currentToken   : %v", currentToken))
		return currentToken, nil
	} else {
		fn := util.TraceStr(0)
		errStr := util.Xs("Out of bounds error %d", reader.position)
		err := SomeError(fn, errStr)
		return NILS, err
	}
}

func (reader *Reader) readAtom() MalType {
	token, _ := reader.peek()
	log.Debug(util.Xs("---------------- atom token: %#v", token))

	intMatches := IntRegex.Match([]byte(token))
	floatMatches := FloatRegex.Match([]byte(token))

	if intMatches {
		log.Debug(util.Xs("intMatches: %#v", intMatches))
		intValue, _ := strconv.ParseInt(token, 10, 64)
		return MalInteger{v: intValue}
	} else if floatMatches {
		log.Debug(util.Xs("floatMatches: %#v", floatMatches))
		floatValue, _ := strconv.ParseFloat(token, 10)
		return MalFloat{v: floatValue}
	} else {
		symbol := MalSymbol{v: token}
		log.Debug("symbol :" + symbol.v)
		return symbol
	}
}

func (reader *Reader) readList() MalType {
	list := MalList{}
	for {
		token, err := reader.next()
		log.Debug(util.Xs("---------------- elem token: %#v", token))

		if &token == nil || token == "" || err != nil {
			if err != nil {
				//log.Warn(err.Error())
			}
			break
		}

		log.Debug("     current : [" + token + "]")
		letter := token[0]

		log.Debug(util.Xs("on (++) counter: %d", reader.counter))
		if letter == '(' {
			reader.counter++
		}

		switch letter {
		case ')':
			log.Debug(util.Xs("on (--) counter: %d", reader.counter))
			if letter == '(' {
				reader.counter--
			}
			break
		default:
			element := reader.readForm()
			if element != nil {
				log.Debug(util.Xs("ELEMENT: %-v", element))
				list.v = append(list.v, element)
			}
		}
	}

	if reader.counter != 0 {
		return MalObject{v: "unbalanced"}
	}
	log.Debug("LIST:" + list.PrintStr())
	return list
}

func (reader *Reader) readForm() MalType {
	token, err := reader.peek()
	log.Debug(util.Xs("--------------- form token: %#v", token))

	if ")" != token {
		if &token == nil || token == "" || err != nil {
			if err != nil {
				//log.Warn(err.Error())
			}
			return nil
		} else {
			switch token[0] {
			case '(':
				reader.counter++
				list := reader.readList()
				log.Debug(util.Xs("LIST: %v", list))
				return list
			default:
				atom := reader.readAtom()
				log.Debug(util.Xs("ATOM: %v", atom))
				return atom
			}
		}
	}
	return nil
}

func CreateReader(tokenStr string) *Reader {
	reader := &Reader{}
	//printFnStart("read_str ", &reader, tokenStr)
	tokens := Tokenize(tokenStr)
	//fmt.Println("===== tokens:", tokens)
	//fmt.Println("===== tokenx:", util.ArrayToString(tokens, "\""))
	reader.tokens = tokens
	reader.position = 0
	reader.counter = 0
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
