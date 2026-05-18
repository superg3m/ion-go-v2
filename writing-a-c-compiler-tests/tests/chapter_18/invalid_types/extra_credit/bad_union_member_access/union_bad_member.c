union s {
    int x;
    int y;
};

union t {
    int blah;
    int y;
};

int main() {
    union s foo = {1};
    return foo.blah; // "union s" has no member "blah"
}