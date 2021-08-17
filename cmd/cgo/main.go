package main

//#cgo CFLAGS: -I../../internal/number
//#cgo LDFLAGS: -L../../internal/number -lnumber
//
//#include "number.h"
import "C"
import "fmt"


// use c static lib
func main() {
	fmt.Println(C.number_add(10, 5))
}
