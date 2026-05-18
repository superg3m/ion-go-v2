// rval in compound expression cannot be void
void f() {
    return;
}

int main() {
    int x = 10;
    x >>= f();
    return 0;
}