// can't apply sizeof directly to a cast expression
// unless the whole cast expression is parenthesized, e.g. sizeof ((char) 1)
int main() {
    return sizeof(char) 1;
}