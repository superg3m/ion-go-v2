// EXPECT: 8
int main() {
	// int x = ---; // uninitialized
	// int x; // zero initialized
	int x = 8;
	int y = x;

	return y;
}