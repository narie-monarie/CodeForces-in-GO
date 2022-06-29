#include <bits/stdc++.h>

#include <algorithm>
#include <iostream>
#include <set>
#include <vector>
using namespace std;

int main() {
	string test = "Hello from Monari";
	set<char> exists;
	for (int i = 0; i < test.length(); i++) {
		char letter = test[i];
		exists.insert(letter);
	}

	for (const auto &n : exists) cout << n << endl;
}

