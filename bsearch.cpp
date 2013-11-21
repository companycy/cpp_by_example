#include <cstdio>

int bad_search(int* array, int n, int v)
{
  int left, right, middle;
  left = 0, right = n;
  while (left < right)		// 
    {
      middle = (left + right) / 2;
      if (array[middle] > v)
        {
	  right = middle - 1;
        }
      else if (array[middle] < v)
        {
	  left = middle + 1;
        }
      else
        {
	  return middle;
        }
    }
  return -1;
}

// some points to avoid missing or dead loop
// http://www.cppblog.com/converse/archive/2009/09/21/96893.html
// mid will always get left+1 in both cases
// if right==n-1
// then
// 1.     if (left<=right)
// 2.        mid = left+1/right-1
// if right==n
// then
// 1.     if (left<right)
// 2.        mid = left+1/right
int search(int* a, int len, int num) {
  int left = 0, right = len-1;
  while (left <= right) {	// can equal since right is len-1
    // int mid = (left+right)/2;
    int mid = left + (right-left)/2; // avoid overflow
    if (a[mid] > num) {
      right = mid-1;
    } else if (a[mid] < num) {
      left = mid+1;
    } else {
      return mid;
    }
  }

  return -1;
}

// it's not consistent compared with previous one
int search2(int* a, int len, int num) {
  int left = 0, right = len;
  while (left < right) {
    int mid = left + (right-left)/2;
    if (a[mid] > num) {
      right = mid;
    } else if (a[mid] < num) {
      left = mid+1;		// !!unconsistent!!
    } else {
      return mid;
    }
  }

  return -1;
}

// recursive version of binary search
int search_recursive(int* a, int left, int right, int num) {
  if (left > right) {
    return -1;
  }
  
  int mid = left+(right-left)/2;
  if (a[mid] > num) {
    return search_recursive(a, left, mid-1, num);
  } else if (a[mid] < num) {
    return search_recursive(a, mid+1, right, num);
  } else {
    return mid;
  }
}

// can't break out while loop
// http://www.cppblog.com/converse/archive/2009/10/05/97905.html
int twoway_search(int* a, int len, int num) {
  int left = -1, right = len;
  int mid;
  while (left+1 != right) {
    mid = left + (right-left)/2;
    if (a[mid] > num) {
      right = mid;
    } else if (a[mid] < num) {
      left = mid;
    }
  }

  if (right >= len || a[right] != num) {
    right = -1;
  }

  return right;
}

int main(int, char**) {
  int a[] = {0, 1, 2, 3, 4, 5};
  int a_len = sizeof(a)/sizeof(int);
  printf("bad found: %d\n", bad_search(a, a_len, 2));
  printf("my found: %d\n", search(a, a_len, 2));
  printf("my found: %d\n", search(a, a_len, 6));
  printf("my found: %d\n", search2(a, a_len, 2));
  printf("my found: %d\n", search2(a, a_len, 6));
  printf("recursive found: %d\n", search_recursive(a, 0, a_len-1, 2));
  printf("recursive found: %d\n", search_recursive(a, 0, a_len-1, 6));
  // printf("my found: %d\n", twoway_search(a, a_len, 2));


  int b[] = {0, 1, 2, 3, 4, 5, 6};
  int b_len = sizeof(b)/sizeof(int);
  printf("bad found: %d\n",  bad_search(b, b_len, 3));
  printf("my found: %d\n", search(b, b_len, 3));
  printf("my found: %d\n", search(b, b_len, -1));
  printf("my found: %d\n", search2(b, b_len, 3));
  printf("my found: %d\n", search2(b, b_len, -1));
  printf("recursive found: %d\n", search_recursive(b, 0, b_len-1, 3)); // !!
  printf("recursive found: %d\n", search_recursive(b, 0, b_len-1, -1)); // !!
  // printf("my found: %d\n",   twoway_search(b, b_len, 3));


  int c[] = {0, 1, 2, 3, 4, 5, 6, 7, 13, 19};
  int c_len = sizeof(c)/sizeof(int);
  printf("recursive found: %d\n",   search_recursive(c, 0, c_len-1, 13));
  printf("recursive found: %d\n",   search_recursive(c, 0, c_len-1, 20));

  return 0;
}
