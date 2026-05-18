// Can't apply postfix ++/-- to void lvalue
extern void *x;

int main() {
    ++(*x)--;
    return 0;
}