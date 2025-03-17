#include "recursive.h"
#include <stdlib.h>
#include <string.h>
#include <omp.h>

static void merge(int arr[], int temp[], int left, int mid, int right) {
    int i = left;     
    int j = mid + 1;  
    int k = left;    

    while (i <= mid && j <= right) {
        if (arr[i] <= arr[j]) {
            temp[k++] = arr[i++];
        } else {
            temp[k++] = arr[j++];
        }
    }

    while (i <= mid) {
        temp[k++] = arr[i++];
    }

    while (j <= right) {
        temp[k++] = arr[j++];
    }

    for (i = left; i <= right; i++) {
        arr[i] = temp[i];
    }
}

static void mergeSortRecursive(int arr[], int temp[], int left, int right) {
    if (left < right) {
        int mid = left + (right - left) / 2; 

        #pragma omp task shared(arr, temp) if(right-left > 10000)
        mergeSortRecursive(arr, temp, left, mid);
        
        #pragma omp task shared(arr, temp) if(right-left > 10000)
        mergeSortRecursive(arr, temp, mid + 1, right);
        
        #pragma omp taskwait
        merge(arr, temp, left, mid, right);
    }
}

void recursiveMergeSort(int arr[], int n) {
    int *temp = (int *)malloc(n * sizeof(int));
    if (temp == NULL) {
        return;
    }

    #pragma omp parallel
    {
        #pragma omp single
        mergeSortRecursive(arr, temp, 0, n - 1);
    }
    
    free(temp);
}