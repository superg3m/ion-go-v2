#ifdef SUPPRESS_WARNINGS
#pragma GCC diagnostic ignored "-Wunused-variable"
#endif
int main() {
    int a;
    {
        int b = a = 1;
    }
    return a;
}