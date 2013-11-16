#include <iostream>
using std::cout;
using std::endl;

// void init(int* a, int len) {
// }

void swap(int* a, int* b) {
  int tmp = *a;
  *a = *b;
  *b = tmp;
}

void myprint(int* a, int left, int right) {
  for (int i = left; i <= right; ++i) {
    cout <<a[i]<<'\t';
  }
  cout<<endl;
}

int partition(int* a, int left, int right) {
  int index = left;
  int pivot = a[left];
  swap((a+left), (a+right));

  // myprint(a, left, right);
  
  for (int i = left; i < right; ++i) {
    if (a[i] > pivot) {
      swap((a+(index++)), (a+i));
    }
  }

  swap((a+right), (a+index));
  // myprint(a, left, right);

  return index;
}

void qsort(int* a, int left, int right) {
  if (left >= right || a == NULL) {
    return;
  }
  
  int index = partition(a, left, right);
  qsort(a, left, index-1);
  qsort(a, index+1, right);
}

int main(int, char**) {
  int a[] = {4, 5, 8, 1, 3, 2};
  myprint(a, 0, 5);
  // init(a, 6);
  qsort(a, 0, 5);
  myprint(a, 0, 5);
  return 0;
}
