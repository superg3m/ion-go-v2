// EXPECT: 3
int foo(int a, int b);

int main() {
    return foo(1, 2);
}

int foo(int a, int b) {
    return a + b;
}