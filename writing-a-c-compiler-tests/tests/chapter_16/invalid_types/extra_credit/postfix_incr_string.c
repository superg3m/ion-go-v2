// can't apply postfix ++/-- to string literals
int main() {
    "foo"++;
    return 0;
}