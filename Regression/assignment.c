// EXPECT: 9
int main() {
	int x = 5;

	if (x < 2) {
		return 9;
	} else {
		return 1;
	}

	return x;
}