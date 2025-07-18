#include <vector>
#include <algorithm>
#include <deque>
#include <iostream>

using namespace std;

// 2,3,5, 8
/*
2 + 2 + 2 + 2
2 + 2 + 3
*/

class Solution
{
public:
    vector<vector<int>> combinationSum(vector<int> &candidates, int target)
    {
        int current_sum{0};
        size_t next_position{0};
        vector<vector<int>> results{};
        deque<int> current_numbers{};

        sort(candidates.begin(), candidates.end());

        // start from poition 0
        dfs(candidates, current_numbers, next_position, current_sum, target, results);

        return results;
    }

    void dfs(vector<int> &candidates, deque<int> &current_numbers, size_t next_pos, int current_sum, int target, vector<vector<int>> &results)
    {
        // for (const auto n : current_numbers)
        // {
        //     cerr << n << " ";
        // }
        // cerr << endl;

        if (current_sum > target)
            return;

        if (current_sum == target)
        {
            vector<int> result{};
            for (const auto n : current_numbers)
            {
                result.push_back(n);
            }
            results.push_back(result);
            return;
        }

        if (next_pos >= candidates.size())
        {
            return;
        }

        // find the number of occurences of next candidate
        auto next_candidate = candidates[next_pos];
        auto curr_pos = next_pos;
        while (next_pos < candidates.size() && next_candidate == candidates[next_pos])
        {
            next_pos += 1;
        }

        // add nothing
        dfs(candidates, current_numbers, next_pos, current_sum, target, results);

        int current_next_pos = curr_pos;
        int thesum = current_sum;
        while (current_next_pos < next_pos)
        {
            next_candidate = candidates[current_next_pos];
            current_numbers.push_back(next_candidate);
            thesum = thesum + next_candidate;
            dfs(candidates, current_numbers, next_pos, thesum, target, results);
            current_next_pos += 1;
        }

        current_next_pos = curr_pos;
        while (current_next_pos < next_pos)
        {
            current_numbers.pop_back();
            current_next_pos += 1;
        }
    }
};

int main()
{
    Solution s{};
    vector<int> candidates{1, 2, 2, 2, 3};
    const auto results = s.combinationSum(candidates, 6);
    if (results.size() == 0)
    {
        cout << "lol";
        return 0;
    }
    for (const auto &result : results)
    {
        cout << "[";
        for (const auto &r : result)
        {
            cout << r << " ";
        }
        cout << "]\n";
    }
}