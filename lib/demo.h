#ifndef DEMO_H_
#define DEMO_H_

#include <iostream>

class Demo {
 public:
  Demo() {}
  ~Demo() {}

  void SayHello() {
    std::cout << "Hello, World!" << std::endl;
  }
};

// give an add function declearation for int


// give an add function declearation for int
int add(int a, int b);
int minus(int a, int b);
#endif // DEMO_H_
