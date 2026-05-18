// can't use string literal as controlling expression in switch statement
int main() {
    switch ("foo") {
        default:
        return 0;
    }
}