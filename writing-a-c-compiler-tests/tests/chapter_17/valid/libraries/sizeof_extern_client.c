
extern double large_array[1000][2000];

int main() {
    return sizeof large_array == 16000000;
}