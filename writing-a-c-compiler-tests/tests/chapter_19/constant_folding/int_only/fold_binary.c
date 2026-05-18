/* Test that we can constant-fold binary expressions, including
 * the arithmetic +, -, *, /, and % operations, and the relational
 * ==, !=, <, <=, >, and >= operations.
 */

// arithmetic tests
int target_add() {
    return 100 + 200;  // 300
}

int target_sub() {
    return 2 - 2147483647;  // -2147483645
}

int target_mult() {
    return 1000 * 1000;  // 1000000
}

int target_div() {
    return 1111 / 4;  // 277
}

int target_rem() {
    return 10 % 3;  // 1
}

// relational tests
int target_eq_true() {
    return 2147483647 == 2147483647;  // 1
}

int target_eq_false() {
    return 2147483647 == 2147483646;  // 0
}

int target_neq_true() {
    return 1111 != 1112;  // 1
}

int target_neq_false() {
    return 1112 != 1112;  // 0
}

int target_gt_true() {
    return 10 > 1;  // 1
}

int target_gt_false() {
    return 10 > 10;  // 0
}

int target_ge_true() {
    return 123456 >= 123456;  // 1
}

int target_ge_false() {
    return 2147 >= 123456;  // 0
}

int target_lt_true() {
    // 256 < 2^30 + 256
    return 256 < 1073742080;  // 1
}

int target_lt_false() {
    return 256 < 0;  // 0
}

int target_le_true() {
    return 123456 <= 123457;  // 1
}

int target_le_false() {
    return 123458 <= 123457;  // 0
}

int val_to_negate = 2147483645;

int main() {
    // arithmetic tests
    if (target_add() != 300) {
        return 1;
    }

    if (target_sub() != -val_to_negate) {
        return 2;
    }
    if (target_mult() != 1000000) {
        return 3;
    }
    if (target_div() != 277) {
        return 4;
    }
    if (target_rem() != 1) {
        return 5;
    }

    // comparisons
    if (target_eq_false()) {
        return 6;
    }
    if (!target_eq_true()) {
        return 7;
    }
    if (target_neq_false()) {
        return 8;
    }
    if (!target_neq_true()) {
        return 9;
    }
    if (target_gt_false()) {
        return 10;
    }
    if (!target_gt_true()) {
        return 11;
    }
    if (target_ge_false()) {
        return 12;
    }
    if (!target_ge_true()) {
        return 13;
    }
    if (target_lt_false()) {
        return 14;
    }
    if (!target_lt_true()) {
        return 15;
    }
    if (target_le_false()) {
        return 16;
    }
    if (!target_le_true()) {
        return 17;
    }
    return 0;  // success
}