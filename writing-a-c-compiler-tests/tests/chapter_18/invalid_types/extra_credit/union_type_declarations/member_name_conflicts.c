/* You can't declare two members of the same union with the same name */
union u {
    int a;
    int a;
};

int main() {
    return 0;
}