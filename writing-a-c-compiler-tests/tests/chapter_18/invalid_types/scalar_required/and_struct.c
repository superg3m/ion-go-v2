struct s {
    int a;
};

int main() {
    struct s x = {1};
    return 0 && x;  // can't apply boolean operators to structs
}