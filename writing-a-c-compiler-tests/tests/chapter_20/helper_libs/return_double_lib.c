// defined in tests/chapter_20/all_types/no_coalescing/return_double.c
double return_double();

// Just use shim to cast return value from double to int
// Expected value is 0.0
int target() {
    return (int) return_double();
}