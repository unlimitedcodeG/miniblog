#include <iostream>
using namespace std;

struct Student {
    char name[10];  // 名字数组长度为10
    int chinese;
    int math;
    int total;
};

int main() {
    Student a[3];  // 创建一个Student数组，大小为3
    for (int j = 0; j < 3; j++) {
        cin >> a[j].name >> a[j].math >> a[j].chinese;
        a[j].total = a[j].math + a[j].chinese;  // 使用加法计算总分
    }

    for (int j = 0; j < 3; j++) {
        cout << a[j].name << " " << a[j].chinese << " " << a[j].math << " "<< a[j].total << endl;
    }
    
    return 0;
}
