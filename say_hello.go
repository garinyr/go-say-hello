package go_say_hello

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) SayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}
