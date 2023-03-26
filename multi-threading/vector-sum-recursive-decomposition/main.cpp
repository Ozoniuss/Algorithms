
/* This program computes the sum of all the elements in a vector by executing a recursive decomposition.
 */

#include <stdio.h>
#include <stdint.h>
#include <memory>
#include <thread>
#include <vector>
#include <future>
#include <chrono>

/*Non-recursive sum on a single thread, to terminate the recursion in a more efficient way*/
int nonrecursiveSum(int const *a, size_t beginIndex, size_t endIndex)
{
    int sum = 0;
    for (size_t i = beginIndex; i < endIndex; ++i)
    {
        sum += a[i];
    }
    return sum;
}

/* Computes the sum of the elements with indexes in [beginIndex, endIndex) by recursive decomposition and using at most nrThreads*/
int recursiveSum(int const *a, size_t beginIndex, size_t endIndex, size_t nrThreads)
{
    if (endIndex == beginIndex + 1)
    {
        return a[beginIndex];
    }

    size_t midPoint = (beginIndex + endIndex) / 2;
    if (nrThreads <= 1)
    {
        return nonrecursiveSum(a, beginIndex, endIndex);
        // return recursiveSum(a, beginIndex, midPoint, 1) + recursiveSum(a, midPoint, endIndex, 1);
    }

    std::future<int> f1 = std::async(std::launch::async, &recursiveSum, a, beginIndex, midPoint, nrThreads / 2);
    int s2 = recursiveSum(a, midPoint, endIndex, nrThreads - (nrThreads / 2));
    return f1.get() + s2;
}

int main(int argc, char **argv)
{
    std::unique_ptr<int[]> a;
    long long nrElements;
    long long nrThreads;
    if (argc != 3 || 1 != sscanf(argv[1], "%lld", &nrElements) || 1 != sscanf(argv[2], "%lld", &nrThreads))
    {
        fprintf(stderr, "Usage: vector_sum_split_work nrElements nrThreads\n");
        return 1;
    }
    a.reset(new int[nrElements]);

    std::chrono::high_resolution_clock::time_point beginTime = std::chrono::high_resolution_clock::now();
    int sum = recursiveSum(a.get(), 0, nrElements, nrThreads);
    std::chrono::high_resolution_clock::time_point endTime = std::chrono::high_resolution_clock::now();
    printf("Used time = %lldus\n", std::chrono::duration_cast<std::chrono::microseconds>(endTime - beginTime).count());
}