// Can't use operands of structure type in bitwise expressions
struct s {int i;};
int main() {
    struct s x = {100};
    int i = 1000;
    x & i;
    return 0;
}