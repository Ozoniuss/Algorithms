#include <iostream>
#include <vector>
#include <unordered_map>
#include <bitset>

using namespace std;
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
class Solution
{
public:
    ListNode *deleteDuplicates(ListNode *head)
    {
        ListNode *shadowHead = new ListNode();
        shadowHead->next = head;
        ListNode *prev = shadowHead;

        ListNode *curr = head;
        ListNode *next = head;

        while (next != nullptr)
        {
            int times_equal = 0;
            while (next->val == curr->val)
            {
                ++times_equal;
                next = next->next;
                if (next == nullptr)
                    break;
            }

            // skip the current node entirely
            if (times_equal > 1)
            {
                prev->next = next;
                curr = next;
            }
            else
            {
                prev = curr;
                curr = next;
            }
        }

        return shadowHead->next;
    }
};
int main()
{
    vector<int> v{1, 2, 2, 2, 3, 3, 4, 5, 5};
    ListNode *l = from_vec(v);
    print_ln(l);
    auto s = Solution{};

    s.deleteDuplicates(l);

    print_ln(l);

    // vector<int> heights{2, 1, 5, 6, 2, 3};
    // cout << s.largestRectangleArea(heights);
}