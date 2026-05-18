/* Label names must be unique within a function */
int main() {
    int x = 0;
label:
    x = 1;
label:
    return 2;
}