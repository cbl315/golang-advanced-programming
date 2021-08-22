package main

import (
	"fmt"
	"unsafe"
)

/*
// struct
struct A {
	int i;
	float f;
	int type;
};
*/
/*
// union
#include <stdint.h>

union B {
    int i;
    float f;
};

union B1 {
    int i;
    float f;
};

union B2 {
    int8_t i8;
    int64_t i64;
};
*/
/*
// enum
enum C {
    ONE,
    TWO,
};
*/
import "C"

func main() {
	var a C.struct_A
	fmt.Println(a.i)
	fmt.Println(a.f)
	fmt.Println(a._type) // _type is A.type, it's key word in go so need to call it with underscore prefix

	var b1 C.union_B1
	fmt.Printf("%T\n", b1) // [4]uint8

	var b2 C.union_B2
	fmt.Printf("%T\n", b2) // [8]uint8

	var b C.union_B
	fmt.Println("b.i:", *(*C.int)(unsafe.Pointer(&b)))
	fmt.Println("b.f:", *(*C.float)(unsafe.Pointer(&b)))

	var c C.enum_C = C.TWO
	fmt.Println(c)
	fmt.Println(C.ONE)
	fmt.Println(C.TWO)
}
