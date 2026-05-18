/* Test that getting the address of a variable does _not_ make that variable
 * live.
 * */

int target() {
    int x = 4;  // initialization is a dead store because we never use the value
                // of x
    int *ptr = &x;
    return ptr == 0;
}

int main() {
    return target();
}