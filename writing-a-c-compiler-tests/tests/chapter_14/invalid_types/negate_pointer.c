/* It's illegal to negate a pointer */
int main() {
    int *x = 0;
    -x;
    return 0;
}