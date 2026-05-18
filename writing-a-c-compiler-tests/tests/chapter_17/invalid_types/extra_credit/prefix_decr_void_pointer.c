// Can't apply prefix or postfix ++/-- to pointers to void
void *malloc(unsigned long size);

int main() {
    void *buff = malloc(100);
    --buff;
    return 0;
}