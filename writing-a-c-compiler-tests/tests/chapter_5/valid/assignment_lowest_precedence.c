#ifdef SUPPRESS_WARNINGS
#ifdef __clang__
#pragma clang diagnostic ignored "-Wconstant-logical-operand"
#endif
#endif

int main() {
    int a;
    a = 0 || 5;
    return a;
}