int get_integer();

// EXPECT: 1
int main() {
    int x = 1;
    int y = 5;
    x = get_integer();
    y = x++;

    return x == 5 && y == 4;
}

int get_integer() {
    return 4;
}
