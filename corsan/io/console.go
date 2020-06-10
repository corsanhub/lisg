package io

import (
	"bufio"
	"fmt"
	"os"
)

func waitForEnterKey(value string) {
	fmt.Printf("%s<Enter>", value)
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
