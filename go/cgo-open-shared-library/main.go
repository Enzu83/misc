package main

/*
#include <dlfcn.h>

typedef void (*run_func_t)(void);

void *open_library(char *name) {
	void *handle = dlopen(name, RTLD_LAZY | RTLD_GLOBAL);
	return handle;
}

void *get_symbol(void *lib_handle, char* symbol_name) {
	return dlsym(lib_handle, symbol_name);
}

void exec_symbol(void *symbol_handle) {
	((run_func_t)symbol_handle)();
}
*/
import "C"

import (
	"fmt"
)

func main() {
	handle := C.open_library(C.CString("./lib/libhello_world.dylib"))
	if handle == nil {
		fmt.Println("Can't find library.")
		return
	}

	symbol_handle := C.get_symbol(handle, C.CString("Run"))
	if symbol_handle == nil {
		fmt.Println("Can't find symbol.")
		return
	}

	C.exec_symbol(symbol_handle)
}