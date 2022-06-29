#include <bits/stdc++.h>

#include <ios>
using namespace std;

int main() {
	ios_base::sync_with_stdio(false);
	cin.tie(NULL);
	int t;
	cin >> t;
	while (t--) {
		int n;
		cin >> n;

		vector<int> a(n);
		for (int i = 0; i < n; i++) {
			cin >> a[i];
		}
		sort(a.begin(), a.end());
		if (n % 2 != 0) {
			int midt = n / 2;
			cout << a[midt] << endl;
		} else {
			int midt = (n / 2) - 1;
			int temp = 0;
			for (int i = 0; i < midt; i++) {
				temp = temp ^ a[i];
				temp = temp ^ a[n - i - 1];
			}
			temp = temp ^ a[midt + 1];
			if (temp == a[midt]) {
				cout << a[midt] << endl;
			} else {
				cout << a[midt] << endl;
			}
		}
	}
	return 0;
}
