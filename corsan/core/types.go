package core

import (
	"bytes"
	"fmt"
)

type MalType interface {
	PrintStr() string
	GetId()
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
	id *string
}

type MalVector struct {
	MalType
	v  []MalType
	id *string
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

func (mal MalInteger) PrintStr() string {
	return fmt.Sprintf("%+v", mal.v)
}

func (mal MalFloat) PrintStr() string {
	return fmt.Sprintf("%+v", mal.v)
}

func (mal MalString) PrintStr() string {
	return fmt.Sprintf("%+v", mal.v)
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
