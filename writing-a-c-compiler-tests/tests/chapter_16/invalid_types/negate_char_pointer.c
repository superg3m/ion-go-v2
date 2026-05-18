/* You can't negate pointers, including pointers to char */
int main() {
    char *x = "foo";
    return -x;
}