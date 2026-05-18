/* Test that we recognize that function calls generate all static variables */
int x = 100;

int get_x() {
    return x;
}

int main() {
    x = 5;  // don't eliminate this!
    int result = get_x();
    x = 10;
    return result;
}