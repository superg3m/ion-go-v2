/* Test that we eliminate the second clause in 1 || x */
int putchar(int c);

int target() {
    return 1 || putchar(97);
}

int main() {
    return target();
}