// EXPECT: 20
int f(int a);

int main() {
    int a = 10;
    // a function declaration is a separate scope,
    // so parameter 'a' doesn't conflict with variable 'a' above

    return f(a);
}

int f(int a) {
    return a * 2;
}