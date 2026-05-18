/* Test constant-folding with INT_MIN */
int target() {
    return -2147483647 - 1;
}

int main() {
    if (~target() != 2147483647) {
        return 1; // fail
    }
    return 0;
}