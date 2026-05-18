/* You can't implicitly convert a char * to a signed char * because they're different types. */
int main() {
    char *c = 0;
    signed char *s = c;
    return (int) s;
}