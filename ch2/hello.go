// +build go1.10

package main

import "fmt"

/*
#include <stdio.h>

static void SayHello(const char* s) {
	puts(s);
}
*/
//#include <hello.h>
//void SayHello5(_GoString_ s);
import "C"

//export SayHello4
func SayHello4(s *C.char) {
	fmt.Print(C.GoString(s))
}

//export SayHello5
func SayHello5(s string) {
	fmt.Print(s)
}

func main() {
	println("hello cgo")
	C.puts(C.CString("hello world\n"))
	C.SayHello(C.CString("hello world, `c` code in go file\n"))
	C.SayHello2(C.CString("hello world, `c` code in c file\n"))
	C.SayHello3(C.CString("hello world, `c` code in cpp file\n"))
	C.SayHello4(C.CString("hello world, `go` code in go file\n"))
	C.SayHello5("hello world, `go` code in go file but using go string\n")
}
