/* A basic test case for eliminating a dead store */

#if defined SUPPRESS_WARNINGS
#pragma GCC diagnostic ignored "-Wunused-variable"
#endif

int target() {
    int x = 10; // this is a dead store
    return 3;
}

int main() {
    return target();
}