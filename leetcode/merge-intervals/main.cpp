#include <iostream>
#include <vector>
#include <unordered_map>
#include <bitset>
#include <string_view>
#include <algorithm>

using namespace std;

#define vi vector<int>
#define vs vector<string>
#define vvi vector<vector<int>>

class Solution
{
public:
    vvi merge(vvi &intervals)
    {
        sort(intervals.begin(), intervals.end());

        vvi ret{};
        vi last{};
        for (const auto &v : intervals)
        {
            cout << v[0] << v[1] << endl;
            if (last.size() == 0)
            {
                last = vi{v};
                continue;
            }
            cout << last[0] << last[1] << endl;

            if (v[0] > last[1])
            {
                ret.push_back(last);
                last = vi{v};
                continue;
            }
            else
            {
                last = vi{last[0], max(v[1], last[1])};
            }
        }
        if (last.size() != 0)
        {
            ret.push_back(last);
        }

        return ret;
    }
};

int main()
{
    cout << __cplusplus << endl;
    vvi v{{2, 6}, {1, 7}, {1, 2}};
    Solution s{};
    const auto ret = s.merge(v);
    for (const auto &v : ret)
    {
        cout << v[0] << " " << v[1] << "\n";
    }
}

/**
 * Definition for singly-linked list.
 */
struct ListNode
{
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

ListNode *from_vec(vector<int> &numbers)
{
    ListNode *head = new ListNode();
    ListNode *curr = head;

    for (const auto num : numbers)
    {
        curr->next = new ListNode(num);
        curr = curr->next;
    }

    return head->next;
}

void print_ln(ListNode *n)
{
    string out;
    ListNode *curr = n;

    out += "[";

    while (curr != nullptr)
    {
        out += to_string(curr->val);
        out += ",";
        curr = curr->next;
    }
    out += "]";
    cout << out << endl;
}

int toInt(const std::string &s)
{
    try
    {
        return std::stoi(s);
    }
    catch (...)
    {
        return -1; // or another sentinel
    }
}
