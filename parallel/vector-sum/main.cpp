/* This program computes the sum of two vectors into a third one, using a configured number of threads.
 * It shows how to split the work evenly between working threads - each thread takes a sequence of consecutive elements.
 */

#include <stdio.h>
#include <stdint.h>
#include <memory>
#include <thread>
#include <vector>
#include <chrono>

/* Does the work for a single thread from computing the pointwise sum of 'a' to 'b' into 'c'*/
void sumForOneThread(int const *a, int const *b, int *c, size_t totalSize, size_t idxThread, size_t nrThreads)
{
    size_t beginIndex = (idxThread * totalSize) / nrThreads;
    size_t endIndex = ((idxThread + 1) * totalSize) / nrThreads;
    for (size_t i = beginIndex; i < endIndex; ++i)
    {
        c[i] = (a[i] + b[i]) % 10;
    }
}

void sumForAllThreads(int const *a, int const *b, int *c, long long nrElems)
{
    for (long long i = 0; i < nrElems; i++)
    {
        c[i] = (a[i] + b[i]) % 10;
    }
}

int main(int argc, char **argv)
{
    std::unique_ptr<int[]> a, b, c;
    long long nrElements;
    long long nrThreads;
    if (argc != 3 || 1 != sscanf(argv[1], "%lld", &nrElements) || 1 != sscanf(argv[2], "%lld", &nrThreads))
    {
        fprintf(stderr, "Usage: vector_sum_split_work nrElements nrThreads\n");
        return 1;
    }
    a.reset(new int[nrElements]);
    b.reset(new int[nrElements]);
    c.reset(new int[nrElements]);

    for (long long i = 0; i < nrElements; i++)
    {
        a[i] = (i % 7) * (i % 11) % 10;
        b[i] = (i % 3) * (i % 13) % 10;
    }

    std::vector<std::thread> threads;
    threads.reserve(nrThreads);

    std::chrono::high_resolution_clock::time_point beginTime = std::chrono::high_resolution_clock::now();
    sumForAllThreads(a.get(), b.get(), c.get(), nrElements);
    std::chrono::high_resolution_clock::time_point endTime = std::chrono::high_resolution_clock::now();
    printf("Used time one thread       = %10lldns\n", std::chrono::duration_cast<std::chrono::nanoseconds>(endTime - beginTime).count());

    beginTime = std::chrono::high_resolution_clock::now();
    for (int i = 0; i < nrThreads; ++i)
    {
        threads.emplace_back(sumForOneThread, a.get(), b.get(), c.get(), nrElements, i, nrThreads);
    }
    for (int i = 0; i < nrThreads; ++i)
    {
        threads[i].join();
    }
    endTime = std::chrono::high_resolution_clock::now();
    printf("Used time multiple threads = %10lldns\n", std::chrono::duration_cast<std::chrono::nanoseconds>(endTime - beginTime).count());
}