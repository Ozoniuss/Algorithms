#include <vector>
#include <algorithm>
#include <iostream>
#include <stdint.h>

using namespace std;

class Solution
{
public:
    int minRemoval(vector<int> &nums, int k)
    {
        if (nums.size() == 0 || nums.size() == 1)
        {
            return 0;
        }
        sort(nums.begin(), nums.end(), less<int>());
        size_t lptr = 0;
        size_t rptr = 0;
        uint64_t diff = 0;

        while (lptr < nums.size())
        {
            while (rptr < nums.size() && uint64_t(nums[lptr] * k) >= uint64_t(nums[rptr]))
            {
                rptr += 1;
            }
            if (uint64_t(rptr) - uint64_t(lptr) > diff)
            {
                diff = uint64_t(rptr) - uint64_t(lptr);
            }
            lptr += 1;
        }

        return int(nums.size()) - diff;
    }
};

int main()
{
    auto s = Solution();
    vector<int> nums{466, 306, 76, 17, 60, 246, 341, 284};
    // vector<int> nums{1, 2, 6, 9};
    // vector<int> nums{2, 1, 5};
    cout << s.minRemoval(nums, 2);
}