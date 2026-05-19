// EXPECT: 1
int main() {
	int x = 5;

	{
		int z = 4;
        return z;
	}

	return x;
}