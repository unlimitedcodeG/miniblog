#include <iostream>
using namespace std;


int main(){
   char ch;
   cin >> ch;

   ch = ((ch>='A')&&(ch<='Z'))?(ch -'A'+ 'a') :ch ;
   cout <<ch; 
   return 0;
}