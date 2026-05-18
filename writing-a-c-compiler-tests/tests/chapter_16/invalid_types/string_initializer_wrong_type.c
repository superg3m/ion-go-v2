/* String literals can only initialize char arrays,
 * not arrays of other types */
int main() {
    long ints[4] = "abc";
    return ints[1];
}