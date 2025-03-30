package main

import "fmt"

// You are playing the Bulls and Cows game with your friend.
//
// You write down a secret number and ask your friend to guess what the number is. When your friend makes a guess, you provide a hint with the following info:
//
// The number of "bulls", which are digits in the guess that are in the correct position.
// The number of "cows", which are digits in the guess that are in your secret number but are located in the wrong position. Specifically, the non-bull digits in the guess that could be rearranged such that they become bulls.
// Given the secret number secret and your friend's guess guess, return the hint for your friend's guess.
//
// The hint should be formatted as "xAyB", where x is the number of bulls and y is the number of cows. Note that both secret and guess may contain duplicate digits.
//
// Example 1:
//
// Input: secret = "1807", guess = "7810"
// Output: "1A3B"
// Explanation: Bulls are connected with a '|' and cows are underlined:
// "1807"
//
//	|
//
// "7810"
// Example 2:
//
// Input: secret = "1123", guess = "0111"
// Output: "1A1B"
// Explanation: Bulls are connected with a '|' and cows are underlined:
// "1123"        "1123"
//
//	|      or     |
//
// "0111"        "0111"
// Note that only one of the two unmatched 1s is counted as a cow since the non-bull digits can only be rearranged to allow one 1 to be a bull.
func getHint(secret string, guess string) string {
	bulls := 0
	cows := 0

	// 각 숫자(0-9)의 등장 횟수를 세기 위한 슬라이스
	secretCount := make([]int, 10)
	guessCount := make([]int, 10)

	// 첫 번째 패스: bulls(소의 머리) 계산
	for i := 0; i < len(secret); i++ {
		s := secret[i]
		g := guess[i]

		if s == g {
			// 같은 위치에 같은 숫자가 있다면 bull
			bulls++
		} else {
			// 다른 위치에 있는 숫자들의 등장 횟수 카운트
			secretCount[s-'0']++
			guessCount[g-'0']++
		}
	}

	// 두 번째 패스: cows(소의 꼬리) 계산
	for i := 0; i < 10; i++ {
		// 각 숫자에 대해 secret과 guess에서 등장한 최소 횟수를 더하면 cows
		cows += min(secretCount[i], guessCount[i])
	}

	// 결과 형식: "xAyB"
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
