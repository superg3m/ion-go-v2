/* Test constant folding of all operations on unsigned longs;
 * make sure that they wrap around correctly,
 * that we evaluate them with unsigned division/comparison functions,
 * and that we can evaluate expressions requiring all 64 bits.
 */
unsigned long target_add() {
    // result exceeds ULONG_MAX and wraps around past 0
    return 18446744073709551615UL + 10ul;
}

unsigned long target_sub() {
    // result is less then 0 and wraps back around past ULONG_MAX
    return 10ul - 12ul;
}

unsigned long target_mult() {
    // wraps back around to 9223372036854775808ul
    return 9223372036854775808ul * 3ul;
}

unsigned long target_div() {
    return 18446744073709551614ul / 10ul;
}

unsigned long target_rem() {
    return 18446744073709551614ul % 10ul;
}

unsigned long target_complement() {
    return ~1ul;
}

unsigned long target_neg() {
    return -(9223372036854775900ul);
}

int target_not() {
    return !4294967296UL;  // 2^32
}

int target_eq() {
    return 18446744073709551615UL == 18446744073709551615UL;
}

int target_neq() {
    // these have identical binary representations except for the most
    // significant bit
    return 9223372036854775809ul != 1ul;
}

int target_gt() {
    // make sure we're using unsigned comparisons;
    // if we interpret these as signed integers,
    // we'll think 9223372036854775809ul is negative and return 0
    return 9223372036854775809ul > 1000ul;
}

int target_ge() {
    // 200ul would be greater if we only considered lower 32 bits
    return 9223372036854775809ul >= 200ul;
}

int target_lt() {
    // as with target_gt, make sure we don't interpret 9223372036854775809ul
    // as a negative signed integer
    return 9223372036854775809ul < 1000ul;
}

int target_le() {
    return 9223372036854775809ul <= 200ul;
}

int target_le2() {
    // make sure we're evaluating <= and not <
    return 9223372036854775809ul <= 9223372036854775809ul;
}

int main() {
    // binary arithmetic
    if (target_add() != 9ul) {
        return 1;
    }
    if (target_sub() != 18446744073709551614ul) {
        return 2;
    }
    if (target_mult() != 9223372036854775808ul) {
        return 3;
    }
    if (target_div() != 1844674407370955161ul) {
        return 4;
    }
    if (target_rem() != 4ul) {
        return 5;
    }

    // unary operators
    if (target_complement() != 18446744073709551614ul) {
        return 6;
    }

    if (target_neg() + 9223372036854775900ul != 0) {
        return 7;
    }

    if (target_not() != 0) {
        return 8;
    }

    // comparisons
    if (!target_eq()) {
        return 9;
    }
    if (!target_neq()) {
        return 10;
    }
    if (!target_gt()) {
        return 11;
    }
    if (!target_ge()) {
        return 12;
    }
    if (target_lt()) {
        return 13;
    }
    if (target_le()) {
        return 14;
    }

    if (!target_le2()) {
        return 15;
    }

    return 0;
}
