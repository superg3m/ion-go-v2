// It's illegal to dereference a label
int main() {
    lbl:
    *lbl;
    return 0;
}