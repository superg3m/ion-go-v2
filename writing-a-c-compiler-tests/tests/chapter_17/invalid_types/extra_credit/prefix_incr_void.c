// Can't apply prefix ++/-- to void lvalue
extern void *x;

int main() {
    ++(*x);
    return 0;
}