// An array that decays to a pointer is not an lvalue, so you can't
// apply ++ or -- to it
int main() {
    int arr[3] = {1, 2, 3};
    arr++;
    return 0;
}