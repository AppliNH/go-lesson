package oop

import (
	"fmt"
)

// A type and its methods should be in the same package
type book struct {
	title string
	price float64
}

type str string

func (b book) printBook() { // copies book, and execute
	fmt.Println(b.title, ": ", b.price, "€")
}

func (b *book) changeBookPrice(newPrice float64) { // doesn't copy book
	b.price = newPrice
}

func (b *book) printBook_bis() { // doesn't copy book
	fmt.Println(b.title, ": ", b.price, "€")
}

func (s *str) getAddr() { // Can't use string as receiver, so I define a new type which is basically the same
	fmt.Println("Address is : ", &s)
}

func ExecLesson2() {

	myBook := book{title: "Shining", price: 15.0}
	myBook.printBook()
	myBook.changeBookPrice(10)
	myBook.printBook_bis()

	var hey str = "salut"
	hey.getAddr() // declaring var hey = "hey" without "str" wouldn't have work, bc hey would have been a string instead of my custom type "str"

}
