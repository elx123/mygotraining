// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare and use function types.
package main

import "fmt"

// event displays global events.
func event(message string) {
	fmt.Println(message)
}

// data is a struct to bind methods to.
type data struct {
	name string
	age  int
}

// event displays events for this data.
func (d *data) event(message string) {
	fmt.Println(d.name, message)
}

// =============================================================================

// fireEvent1 uses an anonymous function type.
func fireEvent1(f func(string)) {
	f("anonymous")
}

// handler represents a function for handling events.
type handler func(string)

// fireEvent2 uses a function type.
func fireEvent2(h handler) {
	h("handler")
}

// =============================================================================

func main() {

	// Declare a variable of type data.
	d := data{
		name: "Bill",
	}

	fireEvent2(d.event)
	d.name = "dsfds"
	fireEvent2(d.event)

}
