void f() { return; }
int main() {
  int i = 0;
  do {
    i = i + 1;

  // void expressions are non-scalar, so they can't be used as controlling conditions
  } while (f());
  return 0;
}