package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	m := greetings.Hello("Gladys")
	fmt.Println(m)
}
