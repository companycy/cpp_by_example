#include <cstdio>


struct list_node {
  int        value;
  list_node* next;
};

list_node* insert(list_node* root, int value) {
  // new a node
  list_node* node = new list_node();
  if (!node) {
    return root;		// return root in case it's updated
  }
  node->value = value;
  node->next = NULL;

  // assume it's ascending
  list_node* p = root, *prev = NULL;
  while (p && p->value < value) {
    prev = p;
    p = p->next;
  }


  if (p) {			// and p->value >= value
    if (p == root) {
      node->next = root;
      root = node;
    } else {
      node->next = p;
      prev->next = node;
    }
  } else {			// p is null, and append to last
    prev->next = node;
  }

  return root;
}

list_node* del(list_node* root, int value) {
  list_node* p = root, *prev = NULL;
  while (p && p->value != value) {
    prev = p;
    p = p->next;
  }

  if (p) {			// found
    if (p == root) {
      p = p->next;
      delete root;		// safe to delete NULL??
      root = p;
    } else if (prev) {		// for safe!!
      prev->next = p->next;
      delete p;
    }
  } else {
    // not found, and do nothing
  }

  return root;
}

void clear(list_node* root) {
  list_node* node = root;
  while (node) {
    list_node* next = node->next;
    delete node;
    node = next;
  }
}

void print(list_node* root) {
  list_node* node = root;
  while (node) {
    printf("%d\t", node->value);
    node = node->next;
  }
  printf("\n");
}

list_node* init() {
  list_node* node = new list_node();
  if (!node) {			// !!!check if it's null
    return NULL;
  }

  node->value = 10;
  node->next = NULL;
  return node;
}

int main(int, char**) {
  list_node* root = init();	// root->value is 10
  if (root) {
    root = insert(root, 13);
    print(root);

    root = insert(root, 18);	// append to last
    print(root);

    root = insert(root, 15);	// insert middle
    print(root);

    root = insert(root, 8);	// insert before root
    print(root);

    root = del(root, 8);    	// del root
    print(root);

    root = del(root, 15);	// del mid
    print(root);

    root = del(root, 18);	// del last node
    print(root);

    root = del(root, 25);	// del non-existing
    print(root);

    root = del(root, 10);
    root = del(root, 13);	// it's now empty list
    print(root);

    clear(root);
    print(root);
  }

  return 0;
}
