struct s;

extern struct s v;

int main() {
  // you can't perform a cast on a struct with incomplete type
  ()v;
  return 0;
}