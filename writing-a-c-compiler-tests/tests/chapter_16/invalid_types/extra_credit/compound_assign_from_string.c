// Can't use a string on the right side of a compound assignment expression
int main() {
    char * s =  "some string ";
    s += "another str";
    return 0;
}