//## 3. (3p) Write a parallel algorithm that computes the product of two big numbers.

#include <iostream>
#include <fstream>
#include <vector>
#include <thread>

using namespace std;

//numbers a and b as vectors
inline vector <int> solve(vector <int> a, vector <int> b, int nr_threads) {
    vector <int> sum;
    int a_1 = a.size();
    int b_1 = b.size();
    int m = b.size() + a.size() - 1;
    sum.resize(m, 0); // it will contain twice the number of elements in a, minus one (so can start with 0)

    vector <thread> thr;

    // add a new thread to the vector of threads
    for (int idx = 0; idx < nr_threads; ++idx) {
        thr.push_back(thread([&, idx]() {
            for (int x = idx; x < m; x += nr_threads) {
                for (int i = 0; i <= x; ++i) {
                    if (i >= a_1 || x - i >= b_1) {
                        continue;
                    }
                    sum[x] += (a[i] * b[x - i]);
                }
            }
            }));
    }   

    //start each vector in the vector of threads
    for (int i = 0; i < thr.size(); ++i) {
        thr[i].join();
    }

    return sum;
}

int main() {
    /*
    ifstream fin("input.in");
    vector <int> a, b;
    int n;
    fin >> n;
    for (int i = 0; i < n; ++i) {
        int x;
        fin >> x;
        a.push_back(x);
    }
    for (int i = 0; i < n; ++i) {
        int x;
        fin >> x;
        b.push_back(x);
    }
    */
    vector <int> a{ 9,2,3};
    vector <int> b{ 9,1,1};
    int head;
    auto sum = solve(a, b, 2);
    for (int i = sum.size() - 1; i >= 1; i--) {
        sum[i - 1] += sum[i] / 10;
        sum[i] = sum[i] % 10;
        if (i == 1) {
            head = sum[0] / 10;
            sum[0] = sum[0] % 10;
        }
    }
    cout << head;
    for (auto it : sum) {
        cout << it;
    }
}
