#include <iostream>
using namespace std;

int main()
{
	int a, b;
	int sum_a = 0, sum_b = 0, dab;

	for (a = 1; a <= 100; a++) sum_a += a*a;
	for (b = 1; b <= 100; b++) {
		sum_b += b;
	}
	sum_b *= sum_b;

	dab = sum_b - sum_a;

	cout << dab;
	return 0;
}