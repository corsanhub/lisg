package core

import (
	"bytes"
	"fmt"
)

type MalType interface {
	PrintStr() string
	Value() interface{}
	Set(interface{})
}

type MalObject struct {
	MalType
	v  string
	id string
}

type MalError struct {
	MalType
	f string
	e string
}

type MalSymbol struct {
	MalType
	v string
}

type MalBoolean struct {
	MalType
	v bool
}

type MalInteger struct {
	MalType
	v int64
}

type MalFloat struct {
	MalType
	v float64
}

type MalString struct {
	MalType
	v string
}

type MalList struct {
	MalType
	v  []MalType
	t  string
	id *string
}

type MalVector struct {
	MalType
	v  []MalType
	id *string
}

type MalMap struct {
	MalType
	v  []MalType
	id *string
}

func GetType(mal MalType) string {
	mType := fmt.Sprintf("%T", mal)
	return mType
}

func (err MalError) Error() string {
	return fmt.Sprintf("MalError thrown in function {%s}: %s", err.f, err.e)
}

func (mal MalObject) PrintStr() string {
	return fmt.Sprintf("%+v", mal.v)
}

func (mal MalSymbol) PrintStr() string {
	return fmt.Sprintf("%+v", mal.v)
}

func (mal MalBoolean) PrintStr() string {
	return fmt.Sprintf("%+v", mal.v)
}

func (mal MalInteger) PrintStr() string {
	return fmt.Sprintf("%+v", mal.v)
}

func (mal MalFloat) PrintStr() string {
	return fmt.Sprintf("%+v", mal.v)
}

func (mal MalString) PrintStr() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("\"%+v\"", mal.v))
	return buffer.String()
}

func (mal MalList) PrintStr() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	for i, item := range mal.v {
		suffix := " "
		if i == (len(mal.v) - 1) {
			suffix = ""
		}
		buffer.WriteString(item.PrintStr() + suffix)
	}
	buffer.WriteString(")")
	return buffer.String()
}

func (mal MalVector) PrintStr() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, item := range mal.v {
		suffix := " "
		if i == (len(mal.v) - 1) {
			suffix = ""
		}
		buffer.WriteString(item.PrintStr() + suffix)
	}
	buffer.WriteString("]")
	return buffer.String()
}

func (mal MalMap) PrintStr() string {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	for i, item := range mal.v {
		suffix := " "
		if i == (len(mal.v) - 1) {
			suffix = ""
		}
		buffer.WriteString(item.PrintStr() + suffix)
	}
	buffer.WriteString("}")
	return buffer.String()
}

func (mal MalObject) Value() interface{} {
	return mal.v
}

func (mal MalSymbol) Value() interface{} {
	return mal.v
}

func (mal MalBoolean) Value() interface{} {
	return mal.v
}

func (mal MalInteger) Value() interface{} {
	return mal.v
}

func (mal MalFloat) Value() interface{} {
	return mal.v
}

func (mal MalString) Value() interface{} {
	return mal.v
}

func (mal MalList) Value() interface{} {
	return mal.v
}

func (mal MalVector) Value() interface{} {
	return mal.v
}

func (mal MalMap) Value() interface{} {
	return mal.v
}

func (mal MalObject) Set(value interface{}) {
	mal.v = value.(string)
}

func (mal MalSymbol) Set(value interface{}) {
	mal.v = value.(string)
}

func (mal MalBoolean) Set(value interface{}) {
	mal.v = value.(bool)
}

func (mal MalInteger) Set(value interface{}) {
	mal.v = value.(int64)
}

func (mal MalFloat) Set(value interface{}) {
	mal.v = value.(float64)
}

func (mal MalString) Set(value interface{}) {
	mal.v = value.(string)
}

func (mal MalList) Set(value interface{}) {
	mal.v = value.([]MalType)
}

func (mal MalVector) Set(value interface{}) {
	mal.v = value.([]MalType)
}

func (mal MalMap) Set(value interface{}) {
	mal.v = value.([]MalType)
}
