/* Function calls kill copies where source value
 * is a variable with static storage duration
 * */

int x = 1;

int f() {
    x = 4;
    return 0;
}

int target() {
    int y = x;  // generate y = x
    f();        // kill y = x
    return y;   // don't
}

int main() {
    if (target() != 1) {
        return 1;
    }

    // make sure x actually got updated
    if (x != 4) {
        return 2;
    }

    return 0;  // success
}