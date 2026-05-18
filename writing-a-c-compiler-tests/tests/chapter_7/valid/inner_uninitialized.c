#ifdef SUPPRESS_WARNINGS
#pragma GCC diagnostic ignored "-Wunused-variable"
#endif
int main() {
    int x = 4;
    {
        int x;
    }
    return x;
}