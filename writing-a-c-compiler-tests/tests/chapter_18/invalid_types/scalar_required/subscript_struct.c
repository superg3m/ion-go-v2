struct s {
  int a;
};

int main() {
  struct s x = {1};
  return x[0]; // can only subscript pointers, not structures
}