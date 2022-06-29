#include <bits/stdc++.h>

#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

void func(int n) {
	if (n > 0) {
		func(n - 1);
		printf("%d \n", n);
	}
}
int main() {
	int x = 3;
	func(x);
}

