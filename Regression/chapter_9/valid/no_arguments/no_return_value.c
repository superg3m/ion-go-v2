// EXPECT: 3
int foo() {
    /* It's legal for a non-void function to not return a value.
     * If the caller tries to use the value of the function, the result is undefined.
     */
    int x = 1;

    return x;
}

int main() {
    /* This is well-defined because we call foo but don't use its return value */
    foo();
    return 3;
}