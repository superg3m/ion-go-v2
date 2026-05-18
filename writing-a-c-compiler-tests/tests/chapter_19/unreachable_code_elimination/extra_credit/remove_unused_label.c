/* Make sure this pass removes unused label instructions */

#ifdef SUPPRESS_WARNINGS
#pragma GCC diagnostic ignored "-Wunused-label"
#endif

int target() {
    lbl:
    return 0;
}

int main() {
    return target();
}