/* This program demonstrates the implementation and the use of a concurrent producer-consumer queue.
 * The main program creates:
 *   - one producer thread that generates numbers from 0 to nrNumbers-1 and writes them in a queue
 *   - a specified number of adders that get the numbers from the producers and add them up. When the
 *       producer closes the queue, the adders write their computed sums into a second queue towards
 *       the printer thread. Note that the adders compete over the numbers from the producer, and it is
 *       unpredictable which adder gets which number.
 *   - finally, one printer thread that gets the partial sums from the adders, adds them all, and prints
 *       the result.
 */

#include <stdio.h>
#include <stdint.h>
#include <thread>
#include <vector>
#include <mutex>
#include <algorithm>
#include <optional>
#include <queue>

template <typename T>
class ProducerConsumerQueue
{
public:
    void enqueue(T const &val)
    {
        std::unique_lock<std::mutex> lck(m_mtx);
        m_queue.push(val);
        m_cv.notify_one();
    }

    /* Gets the next element from the queue. If the queue is empty but not closed, this function blocks.
    If the queue is empty and closed, returns an empty optional.*/
    std::optional<T> dequeue()
    {
        std::optional<T> ret;
        std::unique_lock<std::mutex> lck(m_mtx);
        while (true)
        {
            if (!m_queue.empty())
            {
                ret = m_queue.front();
                m_queue.pop();
                return ret;
            }
            if (m_isEnd)
            {
                return ret;
            }
            m_cv.wait(lck);
        }
    }

    /* Closes the queue.*/
    void close()
    {
        std::unique_lock<std::mutex> lck(m_mtx);
        m_isEnd = true;
        m_cv.notify_all();
    }

private:
    std::mutex m_mtx;
    std::queue<T> m_queue;
    std::condition_variable m_cv;
    bool m_isEnd = false;
};

void producer(ProducerConsumerQueue<int> *pQueue, long long nrNumbers)
{
    for (long long i = 0; i < nrNumbers; ++i)
    {
        pQueue->enqueue(int(i));
    }
    pQueue->close();
}

void adder(ProducerConsumerQueue<int> *pInQueue, ProducerConsumerQueue<int> *pOutQueue)
{
    int sum = 0;
    while (true)
    {
        std::optional<int> x = pInQueue->dequeue();
        if (!x.has_value())
        {
            break;
        }
        sum += *x;
    }
    pOutQueue->enqueue(sum);
}

void printer(ProducerConsumerQueue<int> *pInQueue, int nrAdders)
{
    int sum = 0;
    for (int i = 0; i < nrAdders; ++i)
    {
        std::optional<int> x = pInQueue->dequeue();
        if (!x.has_value())
        {
            fprintf(stderr, "Unexpected end of second queue.\n");
            return;
        }
        sum += *x;
    }
    printf("Sum=%d\n", sum);
}

int main(int argc, char **argv)
{
    int nrAdders;
    long long nrNumbers;
    if (argc != 3 || 1 != sscanf(argv[1], "%d", &nrAdders) || 1 != sscanf(argv[2], "%lld", &nrNumbers))
    {
        fprintf(stderr, "Usage: vector_sum_multithread nrAdders nrNumbers\n");
        return 1;
    }

    std::vector<std::thread> threads;
    threads.reserve(nrAdders + 2);
    ProducerConsumerQueue<int> firstQueue;
    ProducerConsumerQueue<int> secondQueue;

    threads.emplace_back(&printer, &secondQueue, nrAdders);
    for (int i = 0; i < nrAdders; ++i)
    {
        threads.emplace_back(&adder, &firstQueue, &secondQueue);
    }
    threads.emplace_back(&producer, &firstQueue, nrNumbers);

    for (int i = 0; i < nrAdders + 2; ++i)
    {
        threads[i].join();
    }
}