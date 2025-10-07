#include <dlfcn.h>
#include <stdio.h>
#include <stdlib.h>

typedef void (*ffi_func_t)(char **);

int main() {
    void *lib = dlopen("libffi.dylib", RTLD_LAZY | RTLD_GLOBAL);
    ffi_func_t func_sym = dlsym(lib, "ffi_func");

    char **str;
    *str = "hello";

    printf("%s\n", *str);

    func_sym(str);

    printf("%s\n", *str);
}
