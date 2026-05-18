static int f() {
    goto x;
    return 0;
    x:
    return 2;
}

int f_caller() {
    return f();
}