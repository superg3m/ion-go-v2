#ifdef SUPPRESS_WARNINGS
#ifndef __clang__
#pragma GCC diagnostic ignored "-Wunused-value"
#endif
#endif

int main() {
    int a = 0;
    0 && (a = 5);
    return a;
}