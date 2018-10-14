#include <stdio.h>
#include <stdlib.h>
#include <time.h>


int array_total() {
    int i, j, total;
    const int size = 20000;
    int *arr = (int*)malloc(size*size*sizeof(int));
    clock_t st, en;
    
    // fill in the array…    
    srand(time(NULL));
    for(i=0; i<size; i++) {
        for(j=0; j<size; j++) {
            arr[i*size + j] = rand() % 2;
        }
    }
    
    // sum the array…    
    total = 0;
    st = clock();
    for(i=0; i<size; i++) {
        for(j=0; j<size; j++) {
            total += arr[i*size + j];      // sum by row
            // total += arr[j*size + i];   // sum by column
        }
    }
    en = clock();
    
    printf("%f seconds\n", (double)(en - st) / CLOCKS_PER_SEC);
    printf("%d\n", total);
    free(arr);
    return 0;
}

int main(int argc, char **argv) {
    array_total();
    return 0;
}