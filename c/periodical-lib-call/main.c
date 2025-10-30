#include <dlfcn.h>
#include <stdio.h>
#include <unistd.h>

#define LIB_NAME "libhello.dylib"
#define SYMBOL_NAME "Hello"

int main(void) {
    void *lib_handle = dlopen(LIB_NAME, RTLD_LAZY | RTLD_GLOBAL);
    if (!lib_handle) {
        printf("failed to open %s\n", LIB_NAME);
        return 1;
    }

    void (*cb_hello)(void) = dlsym(lib_handle, SYMBOL_NAME);
    if (!cb_hello) {
        printf("failed to get symbol %s\n", SYMBOL_NAME);
        return 1;
    }

    while (true) {
        cb_hello();
        sleep(1);
    }
}
