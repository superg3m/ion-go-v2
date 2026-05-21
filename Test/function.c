int get_integer(int a, int b, int c, int d, int e, int f, int g, int h);

// EXPECT: 1
int main() {
    int x = 1;
    int y = 5;
    x = get_integer(3, 10, 1, 2, 3, 4, 5, 6);
    y = x++;

    return x == 12 && y == 11;
}


int get_integer(int a, int b, int c, int d, int e, int f, int g, int h) {
    return a + g + h;
}
