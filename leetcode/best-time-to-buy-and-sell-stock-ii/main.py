
class Solution:
    def maxProfit(self, prices: list[int]) -> int:
        differences = [prices[i+1]-prices[i] for i in range(len(prices)-1) if prices[i+1] - prices[i] > 0]
        return sum(differences)