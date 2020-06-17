package core

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"corsanhub.com/lisg/corsan/util"
)

const Replacer = "⚾️⚾️⚾️"

func (reader *Reader) readAtom() (MalType, error) {
	token, _ := reader.peek()
	log.Debug(util.Xs("##-------------- --- atom   id: %v, %#v", token, *token))
	if match, _ := regexp.MatchString(`^(true|false)$`, *token); match {
		value, _ := strconv.ParseBool(*token)
		mal := MalBoolean{v: value}
		log.Debug(util.Xs("##-------------- ret bool   id: %v, %#v", token, mal))
		return mal, nil
	} else if match, _ := regexp.MatchString(`^[-|+]{0,1}\d+$`, *token); match {
		value, _ := strconv.ParseInt(*token, 10, 64)
		mal := MalInteger{v: value}
		log.Debug(util.Xs("##-------------- ret aint   id: %v, %#v", token, mal))
		return mal, nil
	} else if match, _ := regexp.MatchString(`^[-|+]{0,1}\d+[.]{1}\d+$`, *token); match {
		value, _ := strconv.ParseFloat(*token, 64)
		mal := MalFloat{v: value}
		log.Debug(util.Xs("##-------------- ret aflt   id: %v, %#v", token, mal))
		return mal, nil
	} else if match, _ := regexp.MatchString("^[^\\s\\[\\]{}('\"`,;)]*$", *token); match {
		mal := MalSymbol{v: *token}
		log.Debug(util.Xs("##-------------- ret symb   id: %v, %#v", token, mal))
		return mal, nil
	} else if match, _ := regexp.MatchString(`^"(?:\\.|[^\\"])*"$`, *token); match {
		str := (*token)[1 : len(*token)-1]
		str = strings.Replace(str, `\\`, Replacer, -1)
		str = strings.Replace(str, `\"`, `"`, -1)
		str = strings.Replace(str, Replacer, "\\", -1)

		malStr := MalString{v: str}
		log.Debug(util.Xs("##-------------- ret atom   id: %v, %#v", token, malStr))
		return malStr, nil
	} else if match, _ := regexp.MatchString(`^.*\s+.*$`, *token); match {
		rchars := `\r\n\t\f\v `
		for pos, char := range rchars {
			fmt.Printf("character %c starts at byte position %d\n", char, pos)
		}
		str := (*token)[1 : len(*token)-1]
		str = strings.Replace(str, `\t`, "", -1)
		str = strings.Replace(str, `\n`, "", -1)
		str = strings.Replace(str, `\r`, "", -1)

		malStr := MalString{v: str}
		log.Debug(util.Xs("##-------------- ret atom   id: %v, %#v", token, malStr))
		return malStr, nil
	} else if match, _ := regexp.MatchString(`^.*\\.*$`, *token); match {
		str := strings.Replace(*token, `\\\\`, "\\", -1)
		malStr := MalString{v: str}
		log.Debug(util.Xs("##-------------- ret atom   id: %v, %#v", token, malStr))
		return malStr, nil
	} else if match, _ := regexp.MatchString(`^.*\s+.*$`, *token); match {
		rchars := `\r\n\t\f\v `
		for pos, char := range rchars {
			fmt.Printf("character %c starts at byte position %d\n", char, pos)
		}
		str := (*token)[1 : len(*token)-1]
		str = strings.Replace(str, `\t`, "", -1)
		str = strings.Replace(str, `\n`, "", -1)
		str = strings.Replace(str, `\r`, "", -1)

		malStr := MalString{v: str}
		log.Debug(util.Xs("##-------------- ret atom   id: %v, %#v", token, malStr))
		return malStr, nil
	} else if match, _ := regexp.MatchString(`^"[^"]*$`, *token); match {
		fn := util.TraceStr(0)
		err := NewError(fn, "unbalanced")

		return nil, err
	} else if (*token)[0] == '"' {
		fn := util.TraceStr(0)
		err := NewError(fn, "expected '\"', got EOF")

		return nil, err
	} else {
		atom := MalObject{v: *token, id: util.RandString(8)}
		log.Debug(util.Xs("##-------------- ret atom   id: %v, %#v", token, atom))
		return atom, nil
	}
}
