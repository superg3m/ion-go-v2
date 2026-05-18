void exit(int status);

struct s;

// you can't define a function with an incomplete return type,
// even if it doesn't actually return a value
struct s return_struct_def() {
  exit(0);
}

int main() { return 0; }