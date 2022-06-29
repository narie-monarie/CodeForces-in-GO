#include <bits/stdc++.h>

#include <iostream>
using namespace std;

int main() {
	ios_base::sync_with_stdio(false);
	cin.tie(NULL);
	int a;
	cin >> a;

	while (a--) {
		int b, c[] = {0}, t = 0;
		cin >> b;
		for (int i = 0; i < b; i++) {
			cin >> c[i];
			t += c[i];
		}
		string result = (t == 0) ? "YES\n" : "NO\n";
		cout << result;
	}
}

