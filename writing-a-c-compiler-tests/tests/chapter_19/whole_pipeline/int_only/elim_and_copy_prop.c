/* If we can replace every use of a variabe with its value,
 * we can delete the copy to that variable
 * */
int target() {
    int x = 10;  // delete this after copy prop rewrites return
    return x;    // rewrite as 'return x' via copy prop
}

int main() {
    return target();
}