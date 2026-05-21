int get_integer(int a, int b, int c, int d, int e, int f, int g, int h);
int putchar(int c);

// EXPECT: 1
int main() {
    int x = 1;
    int y = 5;
    x = get_integer(3, 10, 1, 2, 3, 4, 5, 6);
    y = x++;

    putchar(72);
    putchar(101);
    putchar(108);
    putchar(108);
    putchar(111);
    putchar(44);
    putchar(32);
    putchar(87);
    putchar(111);
    putchar(114);
    putchar(108);
    putchar(100);
    putchar(33);
    putchar(10);

    return x == 15 && y == 14;
}


int get_integer(int a, int b, int c, int d, int e, int f, int g, int h) {
    return a + g + h;
}
