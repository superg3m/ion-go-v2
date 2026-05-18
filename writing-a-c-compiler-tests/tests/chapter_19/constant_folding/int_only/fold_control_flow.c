/* Test constant folding of JumpIfZero and JumpIfNotZero instructions
 * resulting from conditional ?: expressions, if statements, and loops.
 * */

int target_if() {
    if (0)
        return 1;
    return 0;
}

int target_if_else_true() {
    if (1) {
        return 2;
    } else {
        return 3;
    }
}

int target_if_else_false() {
    if (0) {
        return 2;
    } else {
        return 3;
    }
}

int target_conditional_true() {
    return 1 ? 2 : 3;
}

int target_conditional_false() {
    return 0 ? 4 : 5;
}

int target_do_loop() {
    int retval = 0;
    do {
        retval = 10;
    } while (0);
    return retval;
}

int target_while_loop_false() {
    int retval = 0;
    while (0) {
        retval = 10;
    }
    return retval;
}

int target_while_loop_true() {
    int retval = 0;
    while (1048576) {  // 1048576 == 2^20
        retval = 10;
        break;
    }
    return retval;
}

int target_for_loop_true() {
    int retval = 0;
    for (int i = 100; 123;) {
        retval = i;
        break;
    }
    return retval;
}

int target_for_loop_false() {
    int retval = 0;
    for (int i = 100; 0;) {
        retval = i;
        break;
    }
    return retval;
}

int main() {
    if (target_if() != 0) {
        return 1;
    }
    if (target_if_else_true() != 2) {
        return 2;
    }
    if (target_if_else_false() != 3) {
        return 3;
    }
    if (target_conditional_true() != 2) {
        return 4;
    }
    if (target_conditional_false() != 5) {
        return 5;
    }
    if (target_do_loop() != 10) {
        return 6;
    }
    if (target_while_loop_false() != 0) {
        return 7;
    }
    if (target_while_loop_true() != 10) {
        return 8;
    }
    if (target_for_loop_true() != 100) {
        return 9;
    }
    if (target_for_loop_false() != 0) {
        return 10;
    }

    return 0;  // success
}