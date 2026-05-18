union u {
    int i;
    union u self; //illegal; incomplete member type
};

int main() {
    return 0;
}