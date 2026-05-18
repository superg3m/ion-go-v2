struct s;

int main() {
  // can't define a local variable (or any variable) with incomplete type
  struct s v;
  return 0;
}