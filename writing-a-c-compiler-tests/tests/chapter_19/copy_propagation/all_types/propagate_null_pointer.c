/* Test that we can propagate 0 between integer and
 * different pointer types
 * */
long *target() {
    int *ptr = 0;
    long *ptr2 = (long *)ptr;
    return ptr2;  // this should be rewritten as 'return 0'
}

int main() {
    long *result = target();
    return (!result);
}