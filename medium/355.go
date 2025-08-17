package main

import (
	"container/heap"
)

type Tweet struct {
	ID          int
	PublishedAt int64
}

type PriorityTweetQueue []*Tweet

func (pq PriorityTweetQueue) Len() int {
	return len(pq)
}

func (pq PriorityTweetQueue) Less(i, j int) bool {
	return pq[i].PublishedAt > pq[j].PublishedAt
}

func (pq PriorityTweetQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityTweetQueue) Push(x interface{}) {
	item := x.(*Tweet)
	*pq = append(*pq, item)
}

func (pq *PriorityTweetQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

type Twitter struct {
	Counter   int64
	Follower  map[int]map[int]bool
	UserTweet map[int][]*Tweet
}

func Constructor() Twitter {
	return Twitter{
		Counter:   0,
		Follower:  make(map[int]map[int]bool),
		UserTweet: make(map[int][]*Tweet),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	tweet := &Tweet{
		ID:          tweetId,
		PublishedAt: this.Counter,
	}
	this.Counter++
	if this.UserTweet[userId] == nil {
		this.UserTweet[userId] = make([]*Tweet, 0)
	}

	this.UserTweet[userId] = append(this.UserTweet[userId], tweet)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	followeeIDs := make([]int, 0)

	if followees, exists := this.Follower[userId]; exists {
		for followeeId := range followees {
			followeeIDs = append(followeeIDs, followeeId)
		}
	}

	followeeIDs = append(followeeIDs, userId)

	feeds := make(PriorityTweetQueue, 0)
	heap.Init(&feeds)

	for _, followeeId := range followeeIDs {
		if tweets, exists := this.UserTweet[followeeId]; exists {
			for _, tweet := range tweets {
				heap.Push(&feeds, tweet)
			}
		}
	}

	latestFeedIds := make([]int, 0)
	for len(latestFeedIds) < 10 && feeds.Len() > 0 {
		tweet := heap.Pop(&feeds).(*Tweet)
		tweetID := tweet.ID
		latestFeedIds = append(latestFeedIds, tweetID)
	}

	return latestFeedIds
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}

	followee := this.Follower[followerId]
	if followee == nil {
		this.Follower[followerId] = make(map[int]bool)
	}
	this.Follower[followerId][followeeId] = true

}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	followee := this.Follower[followerId]
	delete(followee, followeeId)
}
