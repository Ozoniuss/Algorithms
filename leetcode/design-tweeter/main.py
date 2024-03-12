from typing import List

class Twitter:

    def __init__(self):
        self.follows = set()
        self.tweets = []

    def postTweet(self, userId: int, tweetId: int) -> None:
        self.tweets.append((tweetId,userId))

    def getNewsFeed(self, userId: int) -> List[int]:
        feed =[]
        for item in reversed(self.tweets):
            tweedId, posterID = item
            if (posterID == userId) or ((userId, posterID) in self.follows):
                feed.append(tweedId)
            if len(feed) == 10:
                break
        return feed


    def follow(self, followerId: int, followeeId: int) -> None:
        self.follows.add((followerId, followeeId))

    def unfollow(self, followerId: int, followeeId: int) -> None:
        self.follows.discard((followerId, followeeId))


# Your Twitter object will be instantiated and called as such:
# obj = Twitter()
# obj.postTweet(userId,tweetId)
# param_2 = obj.getNewsFeed(userId)
# obj.follow(followerId,followeeId)
# obj.unfollow(followerId,followeeId)
    
twitter = Twitter()
twitter.postTweet(1, 5)  # User 1 posts a new tweet (id = 5).
print(twitter.getNewsFeed(1))  # User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
twitter.follow(1, 2)     # User 1 follows user 2.
twitter.postTweet(2, 6)  # User 2 posts a new tweet (id = 6).
print(twitter.getNewsFeed(1))  # User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
twitter.unfollow(1, 2);  # User 1 unfollows user 2.
print(twitter.getNewsFeed(1))  # User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.