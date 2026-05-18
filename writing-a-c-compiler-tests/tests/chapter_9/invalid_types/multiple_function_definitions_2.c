/* Function 'foo' is defined twice */
int foo(){
    return 3;
}

int main() {
    // after seeing this declaration, we should still remember that
    // foo was defined earlier
    int foo();
    return foo();
}

int foo(){
    return 4;
}