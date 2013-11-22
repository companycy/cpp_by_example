#include <cstdio>

const int RED = 0;
const int BLACK = 1;

struct rb_node {
  unsigned long  __rb_parent_color;
  rb_node *parent;
  rb_node *right;
  rb_node *left;
  int      key;
  int      color;
};



/*
from algorithm:
left_rotate(T, x):
    y = right[x]
    
    right[x] = left[y]
    p[left[y]] = x
    if p[x] == null
        y = root[T]
    else if x == left[p[node]]
        left[p[x]] = y
    else
        right[p[x]] = y

    left[y] = x
    p[x] = y
*/
/*
       x                    y
     /    \               /   \
    a       y     =>     x     c
          /   \         /  \
         b     c       a    b

*/
// for rotate, the sequence is important!!
// http://blog.csdn.net/v_JULY_v/article/details/6114226
rb_node* rb_rotate_left(rb_node* root, rb_node* node) {
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
right_rotate(T, x):
    y = left[x]

    p[right[y]] = x
    left[x] = right[y]

    if p[x] == null
        root = y
    else if x == left[p[x]]:
        left[p[x]] = y
    else
        right[p[x]] = y

    p[y] = p[x]
    p[x] = y
*/
/*
        x                        y
      /    \                   /    \
      y     c      =>         a      x
    /  \                            /  \
   a    b                          b    c

*/
rb_node* rb_rotate_right(rb_node* root, rb_node* node) {
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

// http://blog.csdn.net/v_JULY_v/article/details/6114226
// code from url uses save last pos searched
rb_node* rb_search_auxiliary(int key, rb_node* root) {
  rb_node* node = root, *parent = NULL;
  while (node && node->key != key) {
    parent = node;
    if (node->key > key) {
      node = node->left;
    } else {
      node = node->right;
    }
  }

  if (node) {			// found
    return node;
  } else {			// not found
    // parent can be last pos searched
    return NULL;
  }
}


/*
while color[p[z]] == RED
    if p[z] = left[p[p[z]]]
        y = right[p[p[z]]]
        if color[y] == RED
	   color[y] = color[p[z]] = BLACK
	   color[p[p[z]]] = RED
	else if z == right[p[z]]
	   z = p[z]
	   left_rotate(T, z)
	   color[p[z]] = BLACK
	   color[p[p[z]]] = RED
	   right_rotate(T, p[p[z]])
	else??

color[root] = BLACK
*/
rb_node* rb_insert_fixup(rb_node* root, rb_node* node) {
  if (!node) {
    return NULL;
  }

  // always assume inserted node is red
  node->color = RED;

  rb_node* parent = node->parent;
  while (parent->color == RED) {
    /*   set color
           gparent                          gparent = red
	  /       \                        /       \
	parent(R)  uncle(R)   =>        parent(B)  uncle(B)
        /                               /  
      node(R)                         node(R)   
    */
    /*    first left rotate parent
            gparent                            gparent
	    /     \                            /     \
       parent(R) uncle(B)     =>            node(R)  uncle(B)
            \                               /
           node(R)                      parent(R)
     */
    /*     set color and right rotate gparent
	     gparent(R)                       node(B)
             /     \                         /       \
	  node(B)  uncle(B)   =>         parent(R)  gparent(R)
	   /                                            \
	parent(R)                                      uncle(B)
     */
    if (parent->parent->left == node->parent) {
      rb_node* uncle = parent->parent->right;
      if (uncle->color == RED) { // node/parent/uncle are all red
	parent->color = BLACK, uncle->color = BLACK;
	parent->parent->color = RED; // set gparent red
	node = parent->parent;	     // always track red node
      } else if (parent->right == node) { // node:red parent: red uncle: black
	node = parent;
	rb_rotate_left(root, node);
	node->parent->color = BLACK;
	node->parent->parent->color = RED;
	rb_rotate_right(root, node->parent->parent);
      }
    } else {
      rb_node* uncle = parent->parent->left;
      if (uncle->color = RED) {
	// same like before
	parent->color = BLACK, uncle->color = BLACK;
	parent->parent->color = RED; // set gparent red
	node = parent->parent;	     // always track red node
      } else if (parent->left == node) {
	node = parent;
	rb_rotate_right(root, node);
	node->parent->color = BLACK;
	node->parent->parent->color = RED;
	// todo(bjcheny): some issue
	rb_rotate_left(root, node->parent->parent);
      }
    }
  }

  root->color = BLACK;
}

/*
  rb_insert(T, z):
     y = null
     x = root(T)

     while x != null
         y = x
	 if key[z] < key[x]
	    x = left[x]
	 else
	    x = right[x]

     p[z] = y

     if y == null
         root(T) = z
     else if key[z] < key[y]
         left[y] = z
     else
         right[y] = z

     left[z] = right[z] = null
     color[z] = red
     rb_insert_fixup(T, z)
*/

rb_node* rb_insert(rb_node* root, rb_node* node) {
  rb_node* p = root, *parent = NULL;
  while (p) {
    parent = p;
    if (p->key > node->key) {
      p = p->left;
    } else {
      p = p->right;
    }
  }

  node->parent = parent;

  if (!parent) {
    root = node;
  } else {
    if (parent->key > node->key) {
      parent->left = node;
    } else {
      parent->right = node;
    }
  }

  node->left = NULL, node->right = NULL;
  node->color = RED;

  rb_insert_fixup(root, node);
}

int main(int, char**) {
  return 0;
}
