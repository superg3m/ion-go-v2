int main() {
  unsigned long x = 0;
  void *v = x; // can't implicitly convert a non-pointer type to a pointer
}