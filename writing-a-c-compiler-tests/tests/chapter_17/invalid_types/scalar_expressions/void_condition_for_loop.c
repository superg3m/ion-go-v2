void foo() {
    return;
}

int main() {
    // void expressions are non-scalar, so they can't be used as controlling conditions
    for (int i = 0; foo(); )
        ;
    return 0;
}