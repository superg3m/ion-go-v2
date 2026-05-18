int main() {
  int x;
  // can't apply sizeof to an expression with incomplete type
  return sizeof(()x);
}