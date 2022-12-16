#include <stdio.h>


int main(void){
    for (int i=0; i < 10000000; i++){
        printf("Parameter value : %d\n", i);
    }
    printf("Here's the end");
    return 0;
}

