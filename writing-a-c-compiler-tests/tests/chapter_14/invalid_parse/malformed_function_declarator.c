/* This is a variation on chapter 9's function_returning_function.c,
 * with a parenthesized declarator. */
int (foo())();

int main() {
    return 0;
}