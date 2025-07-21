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
    ListNode *rotateRight(ListNode *head, int k)
    {
        if (head == nullptr)
            return nullptr;

        ListNode *curr = head;
        int count = 0;
        while (curr != nullptr)
        {
            ++count;
            curr = curr->next;
        }

        k = k % count;
        if (k == 0)
            return head;

        int pos_to_change = count - k - 1;

        int curr_pos = 0;
        curr = head;
        while (curr_pos < pos_to_change)
        {
            curr = curr->next;
            ++curr_pos;
        }

        ListNode *next_node = curr->next;
        ListNode *ret = next_node;
        curr->next = nullptr;

        while (1)
        {
            if (next_node->next == nullptr)
            {
                next_node->next = head;
                break;
            }
            next_node = next_node->next;
        }

        return ret;
    }
};
int main()
{
    vector<int> v{1, 2, 3, 4, 5};
    ListNode *l = from_vec(v);
    print_ln(l);
    auto s = Solution{};

    ListNode *t = s.rotateRight(l, 1);
    print_ln(t);

    // vector<int> heights{2, 1, 5, 6, 2, 3};
    // cout << s.largestRectangleArea(heights);
}