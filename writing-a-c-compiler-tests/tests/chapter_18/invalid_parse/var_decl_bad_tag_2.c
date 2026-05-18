struct s {
    int y;
};

int main() {
    // can't parenthesize struct tag
    struct(s) var;

    return 0;
}