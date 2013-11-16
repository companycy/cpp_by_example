#include <cstdio>

int Left(int i) {
  return 2*i+1;			// start from 0
}

int Right(int i) {
  return 2*i+2;
}

void swap(int* a, int *b) {
  int tmp = *a;
  *a = *b;
  *b = tmp;
}

void print(int a[], int heap_size)
{
  for(int i = 0; i < heap_size; i++) {
    printf("%d ", a[i]);
  }
  printf("\n");
}

void Max_Heapify(int* a, int i, int heap_size) {
  if (i >= heap_size) {
    return;
  }

  int max = i, left = Left(i), right = Right(i);
  // !!left < size || right < size
  if (left < heap_size && a[max] < a[left]) {
    max = left;
  }
  if (right < heap_size && a[max] < a[right]) {
    max = right;
  }

  if (max != i) {
    swap(a+i, a+max);
    Max_Heapify(a, max, heap_size);
  }

}

void Build_Max_Heap(int* a, int heap_size) {
  for (int i = heap_size/2; i >= 0; --i) {
    Max_Heapify(a, i, heap_size);
  }
}

void Heapsort(int* a, int heap_size) {
  for (int i = heap_size-1; i >= 0; --i) {
    swap(a+0, a+i);		// move max to end of array
    Max_Heapify(a, 0, i);
  }
}

void Delete(int* a, int* heap_size) {
  a[0] = a[*heap_size-1];
  --(*heap_size);
  Max_Heapify(a, 0, *heap_size);
}

void InsertIntoHeap(int* a, int num, int* heap_size) {
  a[*heap_size] = num;		// append to last pos
  ++(*heap_size);		// increase heap_size;
  int i = (*heap_size-1), j = (*heap_size)/2-1;
  while (j >= 0) {
    if (*(a+j) < num) {
      a[i] = a[j];
      i = j;
      j = j/2-1;
    } else {
      break;
    }
  }
}

int main(int, char**) {
  int a[] = {3, 4, 9, 10, 5, 2, 7};
  int heap_size = sizeof(a)/sizeof(int);
  Build_Max_Heap(a, heap_size);
  print(a, heap_size);
  // Heapsort(a, heap_size);
  
  print(a, heap_size);
  return 0;
}
