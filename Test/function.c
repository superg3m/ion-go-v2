int get_integer(int z);

// EXPECT: 1
int main() {
    int x = 1;
    int y = 5;
    x = get_integer(5);
    y = x++;

    return x == 10 && y == 9;
}


int get_integer(int z) {
    return z + 4;
}
