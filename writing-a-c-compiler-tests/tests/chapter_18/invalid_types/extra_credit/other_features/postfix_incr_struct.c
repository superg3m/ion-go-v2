// Can't apply prefix or postfix ++/-- to structures
struct s {
    int i;
};

int main() {
    struct s my_struct = {1};
    my_struct++;
    return 0;
}