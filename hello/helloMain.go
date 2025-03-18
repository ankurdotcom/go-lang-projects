package main

import (
	"fmt"
	"log"

	"github.com/ankurdotcom/go-lang-projects/greeting"
)

func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.

	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greeting.Hello("bhavya")

	// If an error was returned, print it to the console and
	// exit the program.

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	//printing multiple messages
	names := []string{"bhavya", "ankur", "saurabh", "aaradhya"}

	// Request greeting mesages for the names.
	messages, err := greeting.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

}
