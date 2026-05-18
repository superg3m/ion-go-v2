/* ++/-- applied to struct members can be a dead store */

struct s {
    int i;
};

int target() {
    struct s my_struct = {4};
    int x = 15;
    my_struct.i--; // dead!
    return x;
}

int main() {
    return target();
}