// You can cast a union to void

union u {
    long l;
    double d;
};

int main() {
    union u x = {1000};
    () x; // just make sure this doesn't cause a type error
    return 0;
}