package main

import (
	"fmt"
	"go-practise/greetings"
	"log"
)

func main() {
	log.SetPrefix("hello-with-greetings: ")
	log.SetFlags(0)

	// msg := greetings.GreetUser("            Anusha            ")
	msg, err := greetings.GreetWithName("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
