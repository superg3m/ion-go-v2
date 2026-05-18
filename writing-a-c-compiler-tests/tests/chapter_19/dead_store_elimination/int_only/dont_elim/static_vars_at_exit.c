/* Make sure we reocgnize that static local variables may be live at exit */
int f() {
    static int i = 10;
    if (i == 5)
        return 0;
    i = 5;  // not a dead store! i is live at exit
    return 1;
}

int main() {

    if (f() != 1) {
        return 1; // fail
    }
    if (f() != 0) {
        return 2; // fail
    }
    return 0;
}