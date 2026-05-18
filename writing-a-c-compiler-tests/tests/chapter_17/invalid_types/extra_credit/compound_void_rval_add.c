// rval in compound expression cannot be void
void f() {
    return;
}

int main() {
    int *x = 0;
    x += f();
    return 0;
}