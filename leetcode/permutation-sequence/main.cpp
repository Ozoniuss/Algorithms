#include <vector>
#include <algorithm>
#include <deque>
#include <iostream>

using namespace std;

class Solution
{
public:
    int fact(int n)
    {
        if (n == 0 || n == 1)
            return 1;
        return n * fact(n - 1);
    }

    // for a number of n digits, n! total permutations
    // permutation on pos k can be written as a * (n-1)! + b * (n-2)! + ...
    // we need to compute a, b, c, ... to find the number
    string getPermutation(int n, int k)
    {
        int vals[10];
        int coef[10];
        int was_no_used[10];
        for (int i = 0; i < n; i++)
        {
            vals[i] = fact(i);
            coef[i] = 0;
            was_no_used[i] = 0;
        }

        int remainder = k;
        int current_factor = n - 1;
        while (current_factor >= 0)
        {
            int d = (remainder - 1) / vals[current_factor];
            remainder = (remainder - 1) % vals[current_factor] + 1;
            coef[current_factor] = d;
            current_factor -= 1;
        }

        string out;

        for (int i = n - 1; i >= 0; i--)
        {
            int digit_used = -1;
            int c = coef[i];

            for (int j = 1; j <= n; j++)
            {
                if (was_no_used[j] == 0)
                    digit_used++;

                if (digit_used == c)
                {
                    was_no_used[j] = 1;
                    out += to_string(j);
                    break;
                }
            }
        }

        return out;
    }
};

int main()
{
    Solution s{};
    // cout << s.getPermutation(4, 9);
    for (int i = 1; i < 24; i++)
    {
        auto out = s.getPermutation(4, i);
        cout << out << "\n";
    }
}