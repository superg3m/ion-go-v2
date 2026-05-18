// Can't apply %= to a double
int main() {
    double d = 5.0;
    d %= 2;
    return (int) d;
}