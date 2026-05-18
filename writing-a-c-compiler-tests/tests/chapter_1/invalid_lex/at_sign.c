/* The @ symbol doesn't appear in any C tokens,
   except inside string or character literals. */
int main() {
    return 0@1;
}