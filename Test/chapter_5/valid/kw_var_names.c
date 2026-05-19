// Test that we can properly lex identifiers that start with keywords
// EXPECT: 5
int main() {
    int return_val = 3;
    int void2 = 2;
    return return_val + void2;
}