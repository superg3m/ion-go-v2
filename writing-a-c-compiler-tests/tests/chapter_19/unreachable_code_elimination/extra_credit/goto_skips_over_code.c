/* Test that we eliminate code that goto jumps over */
int callee() {
    return 1;
}

int target() {
    int x = 10;
    goto end;
    x = callee(); // eliminate this
    end:
    return x;
}

int main() {
    return target();
}