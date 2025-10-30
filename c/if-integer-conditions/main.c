#include <stdio.h>

void check_if(int code) {
    if (!code) {
        printf("%d\n", code);
    }
}

int main(void) {
    check_if(0);
    check_if(1);
    check_if(2);
    
    return 0;
}