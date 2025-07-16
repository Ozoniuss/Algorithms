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
        vector<vector<int>> results{};
        deque<int> current_numbers{};

        dfs(candidates, current_numbers, current_sum, target, results);

        return results;
    }

    void dfs(vector<int> &candidates, deque<int> &current_numbers, int &current_sum, int target, vector<vector<int>> &results)
    {
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

        for (const auto c : candidates)
        {
            if (current_numbers.size() > 0 && current_numbers.back() > c)
            {
                continue;
            }

            current_numbers.push_back(c);
            current_sum += c;
            dfs(candidates, current_numbers, current_sum, target, results);
            current_sum -= c;
            current_numbers.pop_back();
        }
    }
};

int main()
{
    Solution s{};
    vector<int> candidates{2, 3, 5};
    const auto results = s.combinationSum(candidates, 8);
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