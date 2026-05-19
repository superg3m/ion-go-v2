// EXPECT: 0
int main() {
    for (int i = 400; i != 101; i = i - 100)
        if (i == 100)
            return 0;


    return 1;
}
