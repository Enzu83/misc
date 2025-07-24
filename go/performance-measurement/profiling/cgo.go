package main

/*
#include <stdlib.h>

void* leaking_function(size_t size) {
    return malloc(size);
}
*/
import "C"
import (
	"fmt"
)

func run_c_leaking_function() {
    size := C.size_t(200 * 1024 * 1024)
    C.leaking_function(size)
    
	fmt.Println()
}