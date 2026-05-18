// Can't perform bitwise operations with void operands
int main() {
    int x = 10;
    int y = 11;
    x & () y;
    return 0;
}