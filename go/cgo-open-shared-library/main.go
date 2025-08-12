package main

/*
#include <dlfcn.h>
#include <stdio.h>

typedef void (*exported_func_t)(void);

static exported_func_t cb_exported_function;
void ExportedFunction(void);

static void store_callback(exported_func_t cb) {
	cb_exported_function = cb;
}

static void set_callback(void) {
	store_callback(ExportedFunction);
}

typedef struct instance_s {
	exported_func_t cb;
} instance_t;

typedef void (*run_func_t)(const instance_t *);

static void *open_library(char *name, const char **error) {
	void *handle = dlopen(name, RTLD_LAZY | RTLD_GLOBAL);
	if (handle == NULL) {
		*error = "Can't find library";
		return NULL;
	}
	return handle;
}

static void *get_symbol(void *lib_handle, char *symbol_name, const char **error) {
	void *sym_handle = dlsym(lib_handle, symbol_name);
	if (sym_handle == NULL) {
		*error = "Can't find symbol";
		return NULL;
	}
	
	return sym_handle;
}

static void exec_symbol(void *symbol_handle) {
	instance_t instance = { cb_exported_function };

	((run_func_t)symbol_handle)(&instance);
}
*/
import "C"

import (
	"fmt"
)

//export ExportedFunction
func ExportedFunction() {
	fmt.Println("Hello from exported function!")
}

func main() {
	C.set_callback();

	var err *C.char

	handle := C.open_library(C.CString("lib/libhello_world.dylib"), &err)
	if err != nil {
		fmt.Println(C.GoString(err))
		return
	}

	symbol_handle := C.get_symbol(handle, C.CString("Run"), &err)
	if err != nil {
		fmt.Println(C.GoString(err))
		return
	}

	C.exec_symbol(symbol_handle)
}