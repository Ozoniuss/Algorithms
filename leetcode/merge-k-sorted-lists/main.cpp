#include <vector>
#include <algorithm>
#include <queue>
#include <iostream>

using namespace std;

// 2,3,5, 8
/*
2 + 2 + 2 + 2
2 + 2 + 3
*/

struct ListNode
{
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution
{
public:
    ListNode *mergeKLists(vector<ListNode *> &lists)
    {
        ListNode *current = new ListNode();
        ListNode *merged_list_head = nullptr;
        priority_queue<pair<int, ListNode *>, vector<pair<int, ListNode *>>, greater<>> pq{};

        for (auto l : lists)
        {
            if (l != nullptr)
                pq.push({l->val, l});
        }

        while (pq.size())
        {

            auto [val, node] = pq.top();
            pq.pop();

            cout << "val " << val;

            if (node == nullptr)
                continue;

            // advance merged_list
            current->next = new ListNode(val);
            current = current->next;

            if (merged_list_head == nullptr)
            {
                merged_list_head = current;
            }

            // advance node
            node = node->next;

            if (node != nullptr)
                pq.push({node->val, node});
        }
        return merged_list_head;
    }
};

int main()
{
    cout << "a" << endl;
    ListNode *l1 = new ListNode(1);
    l1->next = new ListNode(3);
    l1->next->next = new ListNode(4);

    ListNode *l2 = new ListNode(1);
    l2->next = new ListNode(4);
    l2->next->next = new ListNode(5);

    vector<ListNode *> v{l1, l2};

    Solution s{};
    auto k = s.mergeKLists(v);
    cout << "[ ";
    do
    {
        cout << k->val << " ";
        k = k->next;
    } while (k != nullptr);
    return 0;
}