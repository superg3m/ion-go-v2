/* Test that we add a terminating null byte to the empty string */
int main() {
    char *empty = "";
    return empty[0];
}