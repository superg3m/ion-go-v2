// EXPECT: 11
int foo();

int main() {
    int foo = 3;
    int bar = 4;
    if (foo + bar > 0) {
        /* Function declaration foo shadows variable foo */
        bar = foo();
    }
    /* Variable foo becomes visible again */
    return foo + bar;
}

int foo() {
    return 8;
}