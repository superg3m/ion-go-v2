// EXPECT: 1
int main() {
    int x = 1;
    int y = 5;
    y = x++;

    return x == 2 && y == 1;
}

int get_integer() {
    return 4;
}
