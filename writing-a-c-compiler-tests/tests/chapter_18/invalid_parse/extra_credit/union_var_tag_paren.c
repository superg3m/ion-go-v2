struct s {
    int y;
};

int main() {
    // can't parenthesize union tag
    union(s) var;

    return 0;
}