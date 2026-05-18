/* Test case where we return NaN after constant folding */

int double_isnan(double d); // defined in tests/chapter_13/helper_libs/nan.c

double target() {
    return 0.0 / 0.0;
}

int main() {
    double nan = target();
    if (double_isnan(nan)) {
        return 0; // success
    }

    return 1; // fail
}