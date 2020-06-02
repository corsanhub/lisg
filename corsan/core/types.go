package core

type MalType interface {
}

type MalObject struct {
	MalType
	value *string
}

type MalInteger struct {
	MalType
	value int64
}

type MalFloat struct {
	MalType
	value float64
}

type MalString struct {
	MalType
	value *string
}

type MalList struct {
	MalType
	elements []*MalType
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
