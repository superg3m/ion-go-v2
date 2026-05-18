
#ifdef SUPPRESS_WARNINGS
#ifdef __clang__
#pragma clang diagnostic ignored "-Wconstant-conversion"
#else
#pragma GCC diagnostic ignored "-Woverflow"
#endif
#endif

char return_char() {
    return 5369233654l;  // this will be truncated to -10
}

signed char return_schar() {
    return 5369233654l;  // this will be truncated to -10
}

unsigned char return_uchar() {
    return 5369233654l;  // this will be truncated to 246
}