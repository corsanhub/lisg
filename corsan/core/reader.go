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

func SomeError(fnName, str string) *MalError {
	return &MalError{f: fnName, e: str}
}

type XReader struct {
	position int
	counter  int
	tokens   []*string
}

var XTokenRegexStr = "[\\s,]*(~@|[\\[\\]{}()'`~^@]|\"(?:\\\\.|[^\\\\\"])*\"?|;.*|[^\\s\\[\\]{}('\"`,;)]*)"
var XTokenRegex = regexp.MustCompile(XTokenRegexStr)

func (reader *XReader) XNext() (*string, error) {
	reader.position++
	if reader.position < len(reader.tokens) {
		token := reader.tokens[reader.position]
		log.Debug(util.Xs("##---------------- 🔜next: %#v, id: %v, position: %d", *token, token, reader.position))
		return token, nil
	} else {
		fn := util.TraceStr(0)
		errStr := util.Xs("Out of bounds error %d", reader.position)
		err := SomeError(fn, errStr)
		return nil, err
	}
}

func (reader *XReader) XPeek() (*string, error) {
	if reader.position < len(reader.tokens) {
		token := reader.tokens[reader.position]
		log.Debug(util.Xs("##---------------- 🔷peek: %#v, id: %v, position: %d", *token, token, reader.position))
		return token, nil
	} else {
		fn := util.TraceStr(0)
		errStr := util.Xs("Out of bounds error %d", reader.position)
		err := SomeError(fn, errStr)
		return nil, err
	}
}

func waitForEnterKey(value string) {
	fmt.Printf("%s<Enter>", value)
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func (reader *XReader) xReadAtom() MalType {
	token, _ := reader.XPeek()
	atom := MalObject{v: *token, id: util.RandString(8)}
	log.Debug(util.Xs("##-------------   New atom : %#v", atom))
	return atom
}

func (reader *XReader) xReadList(id *string) MalType {
	list := MalList{id: id}
	log.Debug(util.Xs("##=============   New list.  id: %v, %#v", id, "list"))

	for {
		token, err := reader.XNext()

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
			form := reader.xReadForm()
			list.v = append(list.v, form)
			//waitForEnterKey("list-form :")
		}
	}
	return nil
}

func (reader *XReader) xReadForm() MalType {
	token, _ := reader.XPeek()
	switch (*token)[0] {
	case '(':
		reader.counter++
		list := reader.xReadList(token)
		//log.Debug(util.Xs("##-------------Return list : %#v", *&list))
		return list
	case ')':
		reader.counter--
		log.Debug("It's the end of a list.")
		return nil
	default:
		atom := reader.xReadAtom()
		log.Debug("It's an atom.")
		return atom
	}
}

func XTokenize(str string) []*string {
	matches := XTokenRegex.FindAllStringSubmatch(str, -1)
	tokens := make([]*string, 0)
	for _, g := range matches {
		tokens = append(tokens, &g[1])
	}
	return tokens
}

func XCreateReader(str string) *XReader {
	tokens := XTokenize(str)
	log.Debug(util.Xs("tokens: %s\n", util.PointersToString(tokens, "'")))
	reader := &XReader{tokens: tokens}
	return reader
}

func XReadStr(str string) MalType {
	if &str != nil {
		reader := XCreateReader(str)
		form := reader.xReadForm()

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
	reader := XCreateReader(str)
	form := reader.xReadForm()
	fmt.Printf("form: %#v\n", form)

	println("------------------------------------------------------------------------------------------------------------------------------")

}
