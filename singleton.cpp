#include <cstdio>

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
