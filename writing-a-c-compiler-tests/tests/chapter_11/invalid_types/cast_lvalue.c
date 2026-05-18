/* The result of a cast expression is not an lvalue */

int main() {
    int i = 0;
    i = (long) i = 10;
    return 0;
}