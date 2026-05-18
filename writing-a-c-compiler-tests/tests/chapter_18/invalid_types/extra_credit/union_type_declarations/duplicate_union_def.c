/* YOu can't declare the same union type twice. */

int main() {
    union u {int a;};
    union u {int a;}; // illegal - duplicate declaration
    return 0;
}