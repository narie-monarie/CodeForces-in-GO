package main

import "fmt"

//This Doesnt run but the same code in C++ Does
/*#include<iostream>
using namespace std;
int main(){
    int a, count = 1;
    cin>>a;
    int arr[a];
    for(int i = 0;i<a;i++){
	cin>>arr[i];
    }
    for(int i = 0; i < a-1; i++){
	if(arr[i] != arr[i+1]){
		count++;
	}
   }
   cout<<count++;

}
*/
func main() {
	var a int
	count := 1
	fmt.Scanln(&a)
	arr := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Scanln(&arr[i])
	}
	for i := 0; i < a-1; i++ {
		if arr[i] != arr[i+1] {
			count++
		}
	}
	fmt.Println(count)
}
