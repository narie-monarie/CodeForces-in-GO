#include <bits/stdc++.h>

#include <iostream>
#include <string>
using namespace std;

int main() {
	// ios_base::sync_with_stdio(false);
	// cin.tie(NULL);
	int a;
	int c = 0;
	cin >> a;
	string b;
	cin >> b;
	for (int i = 0; i < a; i++) {
		if (b[i] == b[i + 1]) {
			c++;
		} else {
			c += 0;
		}
	}
	cout << c;
}

