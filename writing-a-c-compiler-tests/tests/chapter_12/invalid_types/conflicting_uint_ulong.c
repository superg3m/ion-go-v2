/* Can't declare the same function with two return types: unsignd int and unsigned long */
unsigned int foo();

unsigned long foo() {
    return 0;
}

int main() {
    return 0;
}