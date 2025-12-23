#include <dlfcn.h>
#include <stdio.h>

int main(void) {
    void *lib = dlopen("/opt/datadog-agent/etc/checks.d/libdatadog-agent-example.dylib", RTLD_NOW | RTLD_LOCAL);
    printf("%s\n", dlerror());
}
