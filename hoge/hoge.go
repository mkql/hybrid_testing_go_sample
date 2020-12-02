package hoge

import "fmt"

// message returns hello
func message() string {
	return "hello"
}

// Do prints "hello" to standard output
func Do() {
	fmt.Println(message())
}