/* Test constant-folding the bitwise &, |, ^, >>, and << expressions with long operands */

long target_and() {
    // 0x0f0f_0f0f_0f0f_0f0f & 0x00ff_00ff_00ff_00ff
    return 1085102592571150095l & 71777214294589695l;
}

long target_or() {
    // 0x0f0f_0f0f_0f0f_0f0f | 0x00ff_00ff_00ff_00ff
    return 1085102592571150095l | 71777214294589695l;
}

long target_xor(){
    // 0x0f0f_0f0f_0f0f_0f0f ^ 0x00ff_00ff_00ff_00ff
    return 1085102592571150095l ^ 71777214294589695l;
}

long target_shift_left() {
    return 1l << 62;
}

long target_shift_right() {
    return 72057589742960640l >> 35;
}

int main() {
    if (target_and() != 4222189076152335l) {
        return 1;
    }

    if (target_or() != 1152657617789587455) {
        return 2;
    }

    if (target_xor() != 1148435428713435120l) {
        return 3;
    }

    if (target_shift_left() != 4611686018427387904l) {
        return 4;
    }

    if (target_shift_right() != 2097151){
        return 5;
    }

     return 0;
}