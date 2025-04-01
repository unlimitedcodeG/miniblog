#include <iostream>
#include <vector>
using namespace std;

int main() {
    int n, d;
    cin >> n >> d;

    vector<int> v(n - 1); // 距离
    for (int i = 0; i < n - 1; ++i) {
        cin >> v[i];
    }

    vector<int> a(n); // 加油价格
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
    }

    int totalCost = 0; // 总花费
    int fuel = 0;      // 当前油量

    for (int i = 0; i < n - 1; ++i) {
        fuel -= v[i]; // 先跑到下一个点，减少油量
        if (fuel < 0) fuel = 0; // 保证不为负

        // 计算之后 d 范围内的最低价
        int minPrice = a[i];
        for (int j = i + 1; j < n && j <= i + d; ++j) {
            if (a[j] < minPrice) {
                minPrice = a[j];
            }
        }

        // 如果当前点是未来 d 范围内最便宜的，就在这里加满（加到 d）
        if (a[i] <= minPrice) {
            int toAdd = d - fuel;
            totalCost += toAdd * a[i];
            fuel += toAdd;
        }
    }

    cout << totalCost << endl;
    return 0;
}
