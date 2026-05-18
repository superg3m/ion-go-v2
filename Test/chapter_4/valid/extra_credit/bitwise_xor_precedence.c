// EXPECT: 1
int main() {
    // ^ has lower precedence than <
    return 5 ^ 7 < 5;
}