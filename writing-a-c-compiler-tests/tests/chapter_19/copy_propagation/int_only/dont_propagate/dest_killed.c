/* Test that updating a variable kills previous
 * copies to that variable
 * */
int foo() {
    return 4;
}

int main() {
    int x = 3;
    x = foo();  // this kills x = 3
    return x;   // don't propagate x = 3
}