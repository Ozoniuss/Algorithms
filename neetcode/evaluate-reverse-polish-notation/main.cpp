#include <iostream>
#include <vector>
#include <string>
#include <cmath>

using namespace std;

int divide(int a, int b)
{
    if (a / static_cast<double>(b) > 0)
    {
        return a / b;
    }
    return static_cast<int>(ceil(a / static_cast<double>(b)));
}

int evalRPM(vector<string> &tokens)
{
    const string signs{"+-/*"};
    vector<string> arguments{};
    for (const string &t : tokens)
    {
        if (signs.find(t) != string::npos)
        {
            const auto arg1 = stoi(arguments[arguments.size() - 2]);
            const auto arg2 = stoi(arguments[arguments.size() - 1]);

            arguments.resize(arguments.size() - 2);

            int val = 0;
            if (t == "+")
            {
                val = arg1 + arg2;
            }
            else if (t == "-")
            {
                val = arg1 - arg2;
            }
            else if (t == "*")
            {
                val = arg1 * arg2;
            }
            else if (t == "/")
            {
                val = divide(arg1, arg2);
            }
            arguments.push_back(to_string(val));
        }
        else
        {
            arguments.push_back(t);
        }
    }
    return stoi(arguments[0]);
}

int main()
{
    vector<string> v{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"};
    cout << evalRPM(v);
}