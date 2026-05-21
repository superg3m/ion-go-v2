int get_integer(int x, int z);

// EXPECT: 1
int main() {
    int x = 1;
    int y = 5;
    x = get_integer(10, 5);
    y = x++;

    return x == 20 && y == 19;
}


int get_integer(int x, int z) {
    return x + z + 4;
}
