#include "iterative.h"
#include <stdlib.h>
#include <string.h>
#include <omp.h>

void merge(int arr[], int temp[], int left, int mid, int right) {
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

void iterativeMergeSort(int arr[], int n) {
    int *temp = (int *)malloc(n * sizeof(int));
    if (temp == NULL) {
        return;  
    }

    for (int size = 1; size < n; size *= 2) {
        #pragma omp parallel for schedule(dynamic) 
        for (int left = 0; left < n; left += 2 * size) {
            int mid = left + size - 1;
            if (mid >= n) mid = n - 1;

            int right = left + 2 * size - 1;
            if (right >= n) right = n - 1;

            int *localTemp = (int *)malloc(n * sizeof(int));
            if (localTemp != NULL) {
                merge(arr, localTemp, left, mid, right);
                free(localTemp);
            }
        }
        
    }

    free(temp);
}