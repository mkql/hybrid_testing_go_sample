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

// ReferToInternalTest refers to internal test helper function
// go build fails
// [main @main]# go build
// # github.com/mkql/hybrid_testing_go_sample/hoge
// hoge/hoge.go:17:2: undefined: someHelperFunc
func ReferToInternalTest() {
	someHelperFunc()
}