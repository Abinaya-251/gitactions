package main

import "fmt"

func main() {
	msg := sayHello("Abinaya.A")
	fmt.Println(msg)
	fmt.Println("Abinaya")
}

func sayHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

