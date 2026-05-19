// EXPECT: 1
int main() {
    for (int a = -2147483647; a % 5 != 0;) {
        a = a + 1;
    }

    return 1;
}
