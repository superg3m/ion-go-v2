// EXPECT: 8
int fib(int n) {
    if (n == 0 || n == 1) {
        return n;
    } else {
        int t1 = fib(n - 1);
        int t2 = fib(n - 2);

        return t1 + t2;
    }
}

int main() {
    int n = 6;
    return fib(n);
}