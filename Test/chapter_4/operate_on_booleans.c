// EXPECT: 0
int main() {
    return ~(0 && 1) - -(4 || 3);
}