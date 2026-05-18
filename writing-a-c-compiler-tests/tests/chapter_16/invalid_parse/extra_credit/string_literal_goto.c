// You can't use a string literal as a label in a goto statement
int main() {
    goto "foo";
    return 0;
}