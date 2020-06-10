package core

type Printer struct {
	position int
	token    string
}

func PrintStr(mal MalType) string {
	return mal.PrintStr()
}

func TestPrinter() {
	//fmt.Printf("Position: %d, Token: %+v\n", currentPosition, currentToken)
	println("Testing printer ...")
	println("---------------------------------------------------------------")

	println("---------------------------------------------------------------")

}
