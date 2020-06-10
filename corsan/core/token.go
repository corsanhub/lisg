package core

import (
	"regexp"

	"corsanhub.com/lisg/corsan/util"
)

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

func Tokenize(str string) []*string {
	matches := XTokenRegex.FindAllStringSubmatch(str, -1)
	tokens := make([]*string, 0)
	for _, g := range matches {
		tokens = append(tokens, &g[1])
	}
	return tokens
}
