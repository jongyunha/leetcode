package main

import (
	"reflect"
	"testing"
)

// Design a simplified version of Twitter where users can post tweets, follow/unfollow another user, and is able to see the 10 most recent tweets in the user's news feed.
//
// Implement the Twitter class:
//
// Twitter() Initializes your twitter object.
// void postTweet(int userId, int tweetId) Composes a new tweet with ID tweetId by the user userId. Each call to this function will be made with a unique tweetId.
// List<Integer> getNewsFeed(int userId) Retrieves the 10 most recent tweet IDs in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user themself. Tweets must be ordered from most recent to least recent.
// void follow(int followerId, int followeeId) The user with ID followerId started following the user with ID followeeId.
// void unfollow(int followerId, int followeeId) The user with ID followerId started unfollowing the user with ID followeeId.
//
// Example 1:
//
// Input
// ["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
// [[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
// Output
// [null, null, [5], null, null, [6, 5], null, [5]]
//
// Explanation
// Twitter twitter = new Twitter();
// twitter.postTweet(1, 5); // User 1 posts a new tweet (id = 5).
// twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
// twitter.follow(1, 2);    // User 1 follows user 2.
// twitter.postTweet(2, 6); // User 2 posts a new tweet (id = 6).
// twitter.getNewsFeed(1);  // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
// twitter.unfollow(1, 2);  // User 1 unfollows user 2.
// twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.
// Test functions
func TestTwitterBasicFunctionality(t *testing.T) {
	twitter := Constructor()

	// Test case from the problem description
	twitter.PostTweet(1, 5)

	// User 1's news feed should return [5]
	newsFeed := twitter.GetNewsFeed(1)
	expected := []int{5}
	if !reflect.DeepEqual(newsFeed, expected) {
		t.Errorf("Expected %v, got %v", expected, newsFeed)
	}

	// User 1 follows user 2
	twitter.Follow(1, 2)

	// User 2 posts a tweet
	twitter.PostTweet(2, 6)

	// User 1's news feed should return [6, 5]
	newsFeed = twitter.GetNewsFeed(1)
	expected = []int{6, 5}
	if !reflect.DeepEqual(newsFeed, expected) {
		t.Errorf("Expected %v, got %v", expected, newsFeed)
	}

	// User 1 unfollows user 2
	twitter.Unfollow(1, 2)

	// User 1's news feed should return [5]
	newsFeed = twitter.GetNewsFeed(1)
	expected = []int{5}
	if !reflect.DeepEqual(newsFeed, expected) {
		t.Errorf("Expected %v, got %v", expected, newsFeed)
	}
}

func TestTwitterMultipleUsers(t *testing.T) {
	twitter := Constructor()

	// Multiple users post tweets
	twitter.PostTweet(1, 1)
	twitter.PostTweet(2, 2)
	twitter.PostTweet(1, 3)
	twitter.PostTweet(2, 4)

	// User 1 follows user 2
	twitter.Follow(1, 2)

	// User 1's news feed should show most recent tweets from both users
	newsFeed := twitter.GetNewsFeed(1)
	expected := []int{4, 3, 2, 1} // Most recent to least recent
	if !reflect.DeepEqual(newsFeed, expected) {
		t.Errorf("Expected %v, got %v", expected, newsFeed)
	}
}

func TestTwitterMoreThan10Tweets(t *testing.T) {
	twitter := Constructor()

	// Post more than 10 tweets
	for i := 1; i <= 15; i++ {
		twitter.PostTweet(1, i)
	}

	// Should only return the 10 most recent
	newsFeed := twitter.GetNewsFeed(1)
	expected := []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6}
	if !reflect.DeepEqual(newsFeed, expected) {
		t.Errorf("Expected %v, got %v", expected, newsFeed)
	}
	if len(newsFeed) != 10 {
		t.Errorf("Expected 10 tweets, got %d", len(newsFeed))
	}
}

func TestTwitterSelfFollow(t *testing.T) {
	twitter := Constructor()

	twitter.PostTweet(1, 1)

	// User tries to follow themselves (should be ignored)
	twitter.Follow(1, 1)

	// Should still see their own tweets
	newsFeed := twitter.GetNewsFeed(1)
	expected := []int{1}
	if !reflect.DeepEqual(newsFeed, expected) {
		t.Errorf("Expected %v, got %v", expected, newsFeed)
	}
}

func TestTwitterUnfollowNonExistentUser(t *testing.T) {
	twitter := Constructor()

	twitter.PostTweet(1, 1)

	// Try to unfollow a user that was never followed
	twitter.Unfollow(1, 2)

	// Should not cause any issues
	newsFeed := twitter.GetNewsFeed(1)
	expected := []int{1}
	if !reflect.DeepEqual(newsFeed, expected) {
		t.Errorf("Expected %v, got %v", expected, newsFeed)
	}
}

func TestTwitterEmptyNewsFeed(t *testing.T) {
	twitter := Constructor()

	// User with no tweets and no following should have empty news feed
	newsFeed := twitter.GetNewsFeed(1)
	expected := []int{}
	if !reflect.DeepEqual(newsFeed, expected) {
		t.Errorf("Expected %v, got %v", expected, newsFeed)
	}
}

func TestTwitterComplexScenario(t *testing.T) {
	twitter := Constructor()

	// Complex scenario with multiple users
	twitter.PostTweet(1, 5)
	twitter.PostTweet(2, 3)
	twitter.PostTweet(1, 101)
	twitter.PostTweet(2, 13)
	twitter.PostTweet(2, 10)
	twitter.PostTweet(1, 2)
	twitter.PostTweet(1, 94)
	twitter.PostTweet(2, 505)
	twitter.PostTweet(1, 333)
	twitter.PostTweet(2, 22)
	twitter.PostTweet(1, 11)
	twitter.PostTweet(1, 205)
	twitter.PostTweet(2, 203)
	twitter.PostTweet(1, 201)
	twitter.PostTweet(2, 213)

	twitter.Follow(1, 2)

	// Should get 10 most recent tweets from both users
	newsFeed := twitter.GetNewsFeed(1)
	expected := []int{213, 201, 203, 205, 11, 22, 333, 505, 94, 2}
	if !reflect.DeepEqual(newsFeed, expected) {
		t.Errorf("Expected %v, got %v", expected, newsFeed)
	}
}

// Benchmark test
func BenchmarkTwitterOperations(b *testing.B) {
	twitter := Constructor()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twitter.PostTweet(i%100, i)
		if i%10 == 0 {
			twitter.Follow(i%100, (i+1)%100)
		}
		if i%20 == 0 {
			twitter.GetNewsFeed(i % 100)
		}
	}
}
