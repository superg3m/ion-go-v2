struct s {
  int a;
};

int main() {
  struct s x = {1};
  // can only apply ~ operator to ints, not structs
  ()~x;
  return 0;
}