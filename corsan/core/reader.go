package core

import "fmt"

type token struct {
	position int
	value    string
	next     *token
}

type Feed struct {
	length int // we'll use it later
	start  *token
}

func (feed *Feed) Append(newtoken *token) {
	if feed.length == 0 {
		feed.start = newtoken
	} else {
		currentPost := feed.start
		for currentPost.next != nil {
			currentPost = currentPost.next
		}
		currentPost.next = newtoken
	}
	feed.length++
}

func DoSomething() {
	println("Doing it ...")

	feed := &Feed{}
	token1 := token{
		position: 0,
		value:    "Lorem ipsum",
	}
	feed.Append(&token1)

	fmt.Printf("Length: %v\n", feed.length)
	fmt.Printf("First: %v\n", feed.start)

	token2 := token{
		value: "Dolor sit amet",
	}
	feed.Append(&token2)

	fmt.Printf("Length: %v\n", feed.length)
	fmt.Printf("First: %v\n", feed.start)
	fmt.Printf("Second: %v\n", feed.start.next)
}
