// Compute the scalar product of two vectors. The vectors must have equal size.

#include <iostream>
#include <vector>
#include <thread>
#include <mutex>

using namespace std;

// The idea here is to sum approximately N / no_threads products into each
// thread, and then add those sums together.
//
// This example adds the sums on position that are congruent modulo no_threads
// in the same thread.
//
// Take the following example:
//
// <1,2,3,4,5,6,7,8,9>
// <1,2,3,4,5,6,7,8,9>
//
// Assuming there are 4 threads, the work would be distributed as follows:
//
// - 1: 1*1 + 5*5 + 9*9
// - 2: 2*2 + 6*6
// - 3: 3*3 + 7*7
// - 4: 4*4 + 8*8
int solveModulo(vector<int> v1, vector<int> v2, int no_threads)
{

    vector<thread> threads;
    int n = v1.size();
    int sum = 0;
    mutex m;
    for (int idx = 0; idx < no_threads; ++idx)
    {
        threads.push_back(thread([v1, v2, idx, n, &sum, &m, no_threads]()
                                 {
            for (int i = idx; i < n; i += no_threads) {
                m.lock(); 
                sum += v1[i] * v2[i];
                m.unlock(); 
            } }));
    }
    for (int i = 0; i < threads.size(); ++i)
    {
        threads[i].join();
    }
    return sum;
}

// This is the same function as the one defined previously, except that the
// elements added by each thread are consecutive. So basically the workload
// is distributed as follows:
//
// - 1: 1*1 + 2*2 + 3*3
// - 2: 4*4 + 5*5
// - 3: 6*6 + 7*7
// - 4: 8*8 + 9*9
int solveConsecutive(vector<int> v1, vector<int> v2, int no_threads)
{

    vector<thread> threads;
    int length = v1.size();
    int sum = 0;
    mutex m;

    int extra = length % no_threads;
    int base_numbers_added = length / no_threads;

    for (int idx = 0; idx < no_threads; ++idx)
    {
        threads.push_back(thread([v1, v2, idx, length, base_numbers_added, extra, &sum, &m]()
                                 {

            // Base number of products to be added, without extra ones
            int steps = base_numbers_added;
            int start = base_numbers_added * idx;

            // Add one of the extra numbers to the first threads
            if (idx < extra) {
                steps++;
                start = (base_numbers_added + 1)*idx;
            } else {
                // Takes into account the extra elements that were added to the
                // previous threads when
                start = start + extra;
            }
            
            for (int i = start; i < start + steps; i ++) {
                m.lock(); 
                sum += v1[i] * v2[i];
                m.unlock(); 
            } }));
    }
    for (int i = 0; i < threads.size(); ++i)
    {
        threads[i].join();
    }
    return sum;
}

int main()
{
    int sum1, sum2;
    sum1 = solveModulo({1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4);
    sum2 = solveConsecutive({1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4);
    cout << sum1 << '\n';
    cout << sum2 << '\n';
}
