package core

import "fmt"

type MalType interface {
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

func (err MalError) Error() string {
	return fmt.Sprintf("MalError thrown in function {%s}: %s", err.f, err.e)
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
	v *string
}

type MalList struct {
	MalType
	v []MalType
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
