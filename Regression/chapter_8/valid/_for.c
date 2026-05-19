// EXPECT: 7
int main() {
	int x = 0;
	int y = 5;
	x = (y += 2);
	while (x < 5) {
		if (x == 3) break;

		x += 1;
	}

	return x;
}