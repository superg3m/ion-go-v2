// void expressions are non-scalar, so they can't be used in logical expressions

void f();
void g();
int main() { return !(1 ? f() : g()); }