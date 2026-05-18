int update_x();

// test that we can use an external variable in a switch statement
int main() {
    update_x(); // set x to 4
    extern int x; // bring x into scope
    switch(x) {
        case 0: return 1; // fail
        case 1: return 2; // fail
        case 4: return 0; // success!
        default: return 4; // fail

    }
}

int x;

int update_x() {
    x = 4;
    return 0;
}
