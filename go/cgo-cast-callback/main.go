package main

/*
#include <stdio.h>
#include <stdbool.h>
#include <string.h>

typedef void (foo_t)(int);

int foo(int arg) {
	printf("foo %d\n", arg);
	return 0;
}

// bar is writing in the memory space pointed by arg2 but **arg2 is read from a garbage memory space
// so calling this function as a `foo_t` function will likely cause a seg fault
int bar(int arg1, char **arg2) {
	printf("bar %d, %s\n", arg1, *arg2);
	*arg2 = strdup("abcdefghi");

	return 0;
}

void *get_foo_cb(void) {
	return (void *)(*foo);
}

void *get_bar_cb(void) {
	return (void *)(*bar);
}

void use_foo_cb(void *foo_cb_untyped, int arg) {
	foo_t *foo_cb = (foo_t *)foo_cb_untyped;
	foo_cb(arg);
}
*/
import "C"

func main() {
	foo_callback := C.get_foo_cb()
	bar_callback := C.get_bar_cb()
	
	C.use_foo_cb(foo_callback, 5)
	C.use_foo_cb(bar_callback, 123456789) // seg fault, data is written in an unpredictable memory space
}
