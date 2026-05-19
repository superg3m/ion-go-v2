// EXPECT: 16
int main() {
    int a = 12345;
    for (int i = 5; i >= 0; i = i - 1)
        a = a / 3;

    return a;
}
