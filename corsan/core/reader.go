package core

import "fmt"

var position int

func init() {
	position = 0
}

func next() Token {
	token := Token{
		value: "some",
	}

	return token
}

func DoSomething() {
	println("Doing something ...")
	token1 := &Token{
		value: "token 1",
	}
	token2 := Token{
		value: "token 2",
		next:  nil,
	}
	token3 := Token{
		value: "token 3",
		next:  nil,
	}

	feed1 := &Feed{
		length: 0,
		start:  token1,
	}
	feed2 := Feed{
		length: 0,
		start:  &token2,
	}
	token1.next = &token2
	token2.next = &token3

	fmt.Printf("feed 1: %+v\n", feed1)
	fmt.Printf("feed 2: %+v\n", feed2)

	fmt.Printf("Start 1: %+v\n", feed1.start)
	fmt.Printf("Start 2: %+v\n", feed2.start)

	fmt.Printf("Next 1: %+v\n", feed1.start.next)
	fmt.Printf("Next 2: %+v\n", feed2.start.next)
}
