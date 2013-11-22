#include <cstdio>


struct rb_node {
  unsigned long  __rb_parent_color;
  rb_node *parent;
  rb_node *right;
  rb_node *left;
};

/*
      node                right
     /    \               /   \
    a     right   =>    node   c
          /   \         /  \
         b     c       a    b

*/
// for rotate, the sequence is important!!
// http://blog.csdn.net/v_JULY_v/article/details/6114226
rb_node* rb_rotate_left(rb_node* node, rb_node* root) {
  rb_node* right = node->right;	   // remember right
  if (node->right = right->left) { // move b(right's left) to node's right
    right->left->parent = node;
  }
  right->left = node;		// move node to right's left

  if (right->parent = node->parent) { // set node's parent to right's parent
    if (node == node->parent->right) { // node is parent's right
      node->parent->right = right;
    } else {
      node->parent->left = right; // node is parent's left
    }
  } else {
    root = right;
  }
  node->parent = right;

  return root;
}

/*
       node                     left
      /    \                   /    \
    left    c      =>         a     node
    /  \                            /  \
   a    b                          b    c

*/
rb_node* rb_rotate_right(rb_node* node, rb_node* root) {
  rb_node* left = node->left;
  if (node->left = left->right) { // mv b(left's right) to node's left
    left->right->parent = node;
  }
  left->right = node;

  if (left->parent = node->parent) { // set node's parent to left's parent
    if (node == node->parent->right) {
      node->parent->right = left;
    } else {
      node->parent->left = left;
    }
  } else {			// parent's null
    root = left;
  }
  node->parent = left;

  return root;
}


int main(int, char**) {
  return 0;
}
