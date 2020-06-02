package core

type MalType interface {
	doSome() string
}

type MalObject struct {
	MalType
	//str string
}

type MalInteger struct {
	MalObject
	value int
}

type MalFloat struct {
	MalObject
	value float32
}

type MalDouble struct {
	MalObject
	value float64
}

type MalString struct {
	MalObject
	value string
}

type MalList struct {
	MalObject
	elements []*MalObject
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
