#include<iostream>
#include<vector>
#include<set>

using namespace std;
int main(){
	vector<int> v;
	v.push_back(3);
	v.push_back(7);
	v.push_back(8);
	for (auto x : v) cout << x << "\n";
	
	set<int> s = {2,5,6,8,8};
	cout << s.size() << "\n"; // 4
	for (auto x : s) cout << x << "\n";
}
