/* Test that we eliminate any code after a return statement. */
int callee() {
    return 1;
}

int target() {
    return 2;

    /* Everything past this point should be optimized away */
    int x = callee();

    if (x) {
        x = 10;
    }

    int y = callee();
    return x + y;
}

int main() {
    return target();
}