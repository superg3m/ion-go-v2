struct s;

void f(struct s param);

extern struct s extern_var;

int main() {
  // can't pass a variable with incomplete type as an argument
  f(extern_var);
}