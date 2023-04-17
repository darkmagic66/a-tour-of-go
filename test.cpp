#include<bits/stdc++.h>

using namespace std;

// pass by value
void swap1(int a, int b){
    int temp;
    temp = a;
    a = b;
    b = temp ;
}

// pass by address
void swap2(int *a, int *b){
    int temp;
    temp = *a;
    *a = *b;
    *b = temp ;
}

// pass by ref
void swap3(int &a, int &b){
    int temp;
    temp = a;
    a = b;
    b = temp ;
}

int main(){
    int a = 2 , b = 3 ;
    // cout << "Hello C++";
    swap2(&a,&b);
    swap3(a,b);
    cout<< a << b ;
    return 0 ;
}