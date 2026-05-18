// void expressions are non-scalar, so they can't be used in logical expressions

int main() {
    return ()1 && 2;
}