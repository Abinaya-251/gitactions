package main

import "fmt"

func main() {
	msg := sayHello("Abinaya")
	fmt.Println(msg)
	fmt.Println("Abinaya")
}

func sayHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
