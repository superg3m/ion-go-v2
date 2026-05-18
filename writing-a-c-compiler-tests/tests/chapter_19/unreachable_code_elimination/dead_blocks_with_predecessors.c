/* Make sure we can eliminate unreachable code even if every unreachable
 * block has a predecessor; in other words, we're traversing the graph to find
 * reachable blocks, not just looking for blocks with no predecessor list.
 * */
#if defined SUPPRESS_WARNINGS
#pragma GCC diagnostic ignored "-Wdiv-by-zero"
#endif

int callee() {
    return 1 / 0;
}

int target() {
    int x = 5;

    return x;

    /* make sure we eliminate this loop even though every block in it has a
     * predecessor */
    for (; x < 10; x = x + 1) {
        x = x + callee();
    }
    return x;
}

int main() {
    return target();
}