#include <iostream>
#include <vector>

using namespace std;

class Solution
{
public:
    vector<int> constructTransformedArray(vector<int> &nums)
    {
        vector<int> ret{};
        const int ln = nums.size();
        for (int i = 0; i < ln; i++)
        {
            cout << (((-10 + 0) % 3) + 3) % 3;
            const int steps = nums[i];
            int m = (((steps + i) % ln) + ln) % ln;
            int val = nums[m];
            cout << val << " " << steps << " " << m << "\n";
            ret.push_back(val);
        }
        return ret;
    }
};

int main()
{
    auto s = Solution{};
    vector<int> input{-10, -10, 4};
    auto ret = s.constructTransformedArray(input);
    cout << ret.size() << "\n";
    for (const auto el : ret)
    {
        cout << el << " ";
    }
}
