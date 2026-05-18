// += and -= operators: RHS must not be a pointer regardless of LHS type
int main() {
    int arr[3] = {1, 2, 3};
    int *elem = arr + 1;
    int i = 0;
    i -= elem;
    return 0;
}