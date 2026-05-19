// EXPECT: 0
int main() {
    int a = 0;
    a = 0 && (a = 5);
    return a;
}