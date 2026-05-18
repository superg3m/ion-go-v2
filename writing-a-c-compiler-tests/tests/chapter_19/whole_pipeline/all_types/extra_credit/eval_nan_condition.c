/* NaN controlling expressions are treated as true */

int target() {
    double nan = 0.0 / 0.0;
    if (nan) {
        return 0; // success
    }
    return 1; // fail
}

int main() {
    return target();
}