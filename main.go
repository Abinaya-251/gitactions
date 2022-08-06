package main

import "fmt"

func main() {
	msg := sayHello("Abinaya")
	fmt.Println(msg)
}

func sayHello(name String) string {
	return fmt.Sprintf("Hi %s", name)
}

