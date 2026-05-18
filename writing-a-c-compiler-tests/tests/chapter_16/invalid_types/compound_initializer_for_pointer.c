/* You can't initialize a pointer with a compound initializer */
int main() {
    char *ptr = {'a', 'b', 'c'};
    return 0;
}