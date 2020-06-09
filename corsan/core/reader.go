package core

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

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

var XTokenRegexStr = "[\\s,]*(~@|[\\[\\]{}()'`~^@]|\"(?:\\\\.|[^\\\\\"])*\"?|;.*|[^\\s\\[\\]{}('\"`,;)]*)"
var XTokenRegex = regexp.MustCompile(XTokenRegexStr)

func (reader *Reader) next() (*string, error) {
	reader.position++
	if reader.position < len(reader.tokens) {
		token := reader.tokens[reader.position]
		log.Debug(util.Xs("##---------------- 🔜next: %#v, id: %v, position: %d", *token, token, reader.position))
		return token, nil
	} else {
		fn := util.TraceStr(0)
		errStr := util.Xs("Out of bounds error %d", reader.position)
		err := NewError(fn, errStr)
		return nil, err
	}
}

func (reader *Reader) peek() (*string, error) {
	if reader.position < len(reader.tokens) {
		token := reader.tokens[reader.position]
		log.Debug(util.Xs("##---------------- 🔷peek: %#v, id: %v, position: %d", *token, token, reader.position))
		return token, nil
	} else {
		fn := util.TraceStr(0)
		errStr := util.Xs("Out of bounds error %d", reader.position)
		err := NewError(fn, errStr)
		return nil, err
	}
}

func waitForEnterKey(value string) {
	fmt.Printf("%s<Enter>", value)
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func (reader *Reader) readAtom() MalType {
	token, _ := reader.peek()
	atom := MalObject{v: *token, id: util.RandString(8)}
	log.Debug(util.Xs("##-------------   New atom : %#v", atom))
	return atom
}

func (reader *Reader) readList(id *string) MalType {
	list := MalList{id: id}
	log.Debug(util.Xs("##=============   New list.  id: %v, %#v", id, "list"))

	for {
		token, err := reader.next()

		if token == nil || *token == "" || err != nil {
			if err != nil {
				//log.Warn(err.Error())
			}
			log.Debug("Breaking 📰 ...")
			break
		}

		switch (*token)[0] {
		case ')':
			reader.counter--
			//waitForEnterKeywaitForEnterKey("list-retr :")
			log.Debug(util.Xs("##=============   Ret list.  id: %v, %#v", id, "list"))
			return list
		default:
			form := reader.readForm()
			list.v = append(list.v, form)
			//waitForEnterKey("list-form :")
		}
	}
	return nil
}

func (reader *Reader) readVector(id *string) MalType {
	vector := MalVector{id: id}
	log.Debug(util.Xs("##=============   New vect.  id: %v, %#v", id, "list"))

	for {
		token, err := reader.next()

		if token == nil || *token == "" || err != nil {
			if err != nil {
				//log.Warn(err.Error())
			}
			log.Debug("Breaking 📰 ...")
			break
		}

		switch (*token)[0] {
		case ']':
			reader.counter--
			//waitForEnterKeywaitForEnterKey("list-retr :")
			log.Debug(util.Xs("##=============   Ret vect.  id: %v, %#v", id, "list"))
			return vector
		default:
			form := reader.readForm()
			vector.v = append(vector.v, form)
			//waitForEnterKey("list-form :")
		}
	}
	return nil
}

func (reader *Reader) readForm() MalType {
	token, _ := reader.peek()
	switch (*token)[0] {
	case '(':
		reader.counter++
		list := reader.readList(token)
		//log.Debug(util.Xs("##-------------Return list : %#v", *&list))
		return list
	case '[':
		reader.counter++
		list := reader.readVector(token)
		//log.Debug(util.Xs("##-------------Return list : %#v", *&list))
		return list
	case ')':
		reader.counter--
		log.Debug("It's the end of a list.")
		return nil
	case ']':
		reader.counter--
		log.Debug("It's the end of a vector.")
		return nil
	default:
		atom := reader.readAtom()
		log.Debug("It's an atom.")
		return atom
	}
}

func Tokenize(str string) []*string {
	matches := XTokenRegex.FindAllStringSubmatch(str, -1)
	tokens := make([]*string, 0)
	for _, g := range matches {
		tokens = append(tokens, &g[1])
	}
	return tokens
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
		form := reader.readForm()

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

	//str = "(a b)"
	str := "(()())"
	reader := CreateReader(str)
	form := reader.readForm()
	fmt.Printf("form: %#v\n", form)

	println("------------------------------------------------------------------------------------------------------------------------------")

}
