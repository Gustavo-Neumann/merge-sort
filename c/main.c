#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <string.h>
#include <omp.h>
#include "iterative.h"
#include "recursive.h"

#define ARRAY_SIZE 10000
#define RUNS 10

int threadCounts[] = {2, 4, 8, 16};
int numThreadCounts = 4;

double iterativeTimes[5][RUNS];
double recursiveTimes[5][RUNS]; 

void generateRandomArray(int arr[], int size) {
    for (int i = 0; i < size; i++) {
        arr[i] = rand() % ARRAY_SIZE;
    }
}

int isSorted(int arr[], int size) {
    for (int i = 1; i < size; i++) {
        if (arr[i-1] > arr[i]) {
            return 0; 
        }
    }
    return 1;
}

int main() {
    srand(time(NULL));
    
    FILE *file = fopen("results.txt", "w");
    if (file == NULL) {
        fprintf(stderr, "Erro ao criar arquivo\n");
        return 1;
    }
    
    for (int t = 0; t < numThreadCounts; t++) {
        int threadCount = threadCounts[t];
        printf("Executando com %d threads...\n", threadCount);
        
        // Para cada execução
        for (int run = 0; run < RUNS; run++) {
            int *arr1 = (int *)malloc(ARRAY_SIZE * sizeof(int));
            int *arr2 = (int *)malloc(ARRAY_SIZE * sizeof(int));
            
            if (arr1 == NULL || arr2 == NULL) {
                fprintf(stderr, "Erro de alocação de memória\n");
                exit(1);
            }
            
            generateRandomArray(arr1, ARRAY_SIZE);
            memcpy(arr2, arr1, ARRAY_SIZE * sizeof(int));
  
            double start_time = omp_get_wtime();
            
            omp_set_num_threads(threadCount);
     
            iterativeMergeSort(arr1, ARRAY_SIZE);
            
            double end_time = omp_get_wtime();
            iterativeTimes[t][run] = (end_time - start_time) * 1000.0;  // em ms
            
            if (!isSorted(arr1, ARRAY_SIZE)) {
                fprintf(stderr, "Aviso: MergeSort iterativo não ordenou o array corretamente!\n");
            }
            
            start_time = omp_get_wtime();

            omp_set_num_threads(threadCount);

            recursiveMergeSort(arr2, ARRAY_SIZE);
            
            end_time = omp_get_wtime();
            recursiveTimes[t][run] = (end_time - start_time) * 1000.0;  // em ms

            if (!isSorted(arr2, ARRAY_SIZE)) {
                fprintf(stderr, "Aviso: MergeSort recursivo não ordenou o array corretamente!\n");
            }
            
            free(arr1);
            free(arr2);
        }

        fprintf(file, "Threads: %d\n", threadCount);
        
        fprintf(file, "Iterative Times (ms):\n");
        double totalIterative = 0.0;
        for (int i = 0; i < RUNS; i++) {
            fprintf(file, "Run %d: %.3f\n", i+1, iterativeTimes[t][i]);
            totalIterative += iterativeTimes[t][i];
        }
        double avgIterative = totalIterative / RUNS;
        fprintf(file, "Avg: %.3f\n\n", avgIterative);
        
        fprintf(file, "Recursive Times (ms):\n");
        double totalRecursive = 0.0;
        for (int i = 0; i < RUNS; i++) {
            fprintf(file, "Run %d: %.3f\n", i+1, recursiveTimes[t][i]);
            totalRecursive += recursiveTimes[t][i];
        }
        double avgRecursive = totalRecursive / RUNS;
        fprintf(file, "Avg: %.3f\n\n", avgRecursive);
        
        fprintf(file, "----------------------------\n\n");
    }
    
    fprintf(file, "Threads\tIterativo (ms)\tRecursivo (ms)\n");
    for (int t = 0; t < numThreadCounts; t++) {
        double totalIterative = 0.0, totalRecursive = 0.0;
        
        for (int i = 0; i < RUNS; i++) {
            totalIterative += iterativeTimes[t][i];
            totalRecursive += recursiveTimes[t][i];
        }
        
        double avgIterative = totalIterative / RUNS;
        double avgRecursive = totalRecursive / RUNS;
        
        fprintf(file, "%d\t%.3f\t\t%.3f\n", threadCounts[t], avgIterative, avgRecursive);
    }
    
    fclose(file);
    printf("Benchmark concluído! Resultados salvos em 'results.txt'\n");
    
    return 0;
}