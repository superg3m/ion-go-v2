// EXPECT: 3
int main() {
	int x = 0;
	while (x < 5) {
		if (x == 3) break;

		x += 1;
	}

	return x;
}