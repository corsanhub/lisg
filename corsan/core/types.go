package core

import (
	"bytes"
	"fmt"
)

type MalType interface {
	PrintStr() string
	Value()
}

type MalObject struct {
	MalType
	v string
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
	v []MalType
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
		log.Debug(fmt.Sprintf("len(mal.v): %+v, i: %d", len(mal.v), i))
		if i == (len(mal.v) - 1) {
			suffix = ""
		}
		buffer.WriteString(item.PrintStr() + suffix)
	}
	buffer.WriteString(")")
	return buffer.String()
}

// func (malObject MalObject) doSome() string {
// 	return malObject.str
// }

// func TestTypes() {
// 	fmt.Println("Testing types ...")

// 	ss := &MalString{MalObject{}, "Some value"}
// 	ss.str = "Daj K'ptin"

// 	fmt.Println("str : ", ss.str)
// 	fmt.Println("some: ", ss.doSome())

// }
