#include <cstdio>

// A macro to disallow the copy constructor and operator= functions
// This should be used in the private: declarations for a class
#define DISALLOW_COPY_AND_ASSIGN(TypeName) \
  TypeName(const TypeName&);               \
  void operator=(const TypeName&)

class Foo {
public:
  Foo(int f);
  ~Foo();

private:
  DISALLOW_COPY_AND_ASSIGN(Foo);
};


class Singleton {
public:
  static Singleton& getInst() {
    static Singleton inst;
    return inst;
  }

private:
  explicit Singleton() {	// private ctor
    printf("s:ctor\n");
  };
  ~Singleton() {
    printf("s:dtor\n");
  };
  Singleton(Singleton&);	    // disable copy ctor
  Singleton& operator=(Singleton&); // disable operator
};

int main(int, char**) {
  Singleton& s = Singleton::getInst(); // can only be reference
  return 0;
}
