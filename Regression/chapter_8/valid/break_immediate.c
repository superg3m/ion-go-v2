// EXPECT: 1
int main() {
    int a = 10;
    while ((a = 1))
        break;
    return a;
}
