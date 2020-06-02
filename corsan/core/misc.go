package core

import (
	"container/list"
	"fmt"
)

func another() {
	// Create a new list and insert elements in it.
	l := list.New()
	l.PushBack(1) // 1
	l.PushBack(2) //  1 -> 2
	l.PushBack(3) //  1 -> 2 -> 3

	// get the head of the list
	ele := l.Front()
	fmt.Println(ele.Value)

	// get next element of the list
	ele = ele.Next()
	fmt.Println(ele.Value)

	// get next element of the list
	ele = ele.Next()
	fmt.Println(ele.Value)
}

// func (feed *Feed) Append(newtoken *Token) {
// 	if feed.length == 0 {
// 		//newtoken.position = 0
// 		feed.start = newtoken
// 		feed.currentPosition = 0
// 	} else {
// 		currentToken := feed.start
// 		for currentToken.next != nil {
// 			currentToken = currentToken.next
// 			feed.currentPosition++
// 		}
// 		currentToken.next = newtoken
// 	}
// 	feed.length++
// }

// func more() {
// 	feed := &Feed{}
// 	token1 := Token{
// 		value: "Lorem ipsum",
// 	}
// 	feed.Append(&token1)

// 	token2 := Token{
// 		value: "Dolor sit amet",
// 	}
// 	feed.Append(&token2)
// 	s := "Hello, 世界"
// 	tokenize(s)

// }

// func andMore() {
// 	token1 := &Token{
// 		value: "token 1",
// 	}
// 	token2 := &Token{
// 		value: "token 2",
// 		next:  nil,
// 	}
// 	token3 := &Token{
// 		value: "token 3",
// 		next:  nil,
// 	}

// 	feed1 := &Feed{
// 		length: 0,
// 		start:  token1,
// 	}
// 	feed2 := Feed{
// 		length: 0,
// 		start:  token2,
// 	}
// 	token1.next = token2
// 	token2.next = token3

// 	fmt.Printf("feed 1: %+v\n", feed1)
// 	fmt.Printf("feed 2: %+v\n", feed2)

// 	fmt.Printf("Start 1: %+v\n", feed1.start)
// 	fmt.Printf("Start 2: %+v\n", feed2.start)

// 	fmt.Printf("Next 1: %+v\n", feed1.start.next)
// 	fmt.Printf("Next 2: %+v\n", feed2.start.next)

// }