package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	// case 1
	if LengthOfLongestSubstring("a") == 1 {
		t.Log("test ok")
	} else {
		t.Fail()
	}

	// case 2
	if LengthOfLongestSubstring("abcabcbb") == 3 {
		t.Log("test ok")
	} else {
		t.Fail()
	}

	// case 3
	if LengthOfLongestSubstring("pwwkew") == 3 {
		t.Log("test ok")
	} else {
		t.Fail()
	}

	// case 4
	if LengthOfLongestSubstring("dvdf") == 3 {
		t.Log("test ok")
	} else {
		t.Fail()
	}

	// case 5
	if LengthOfLongestSubstring("ckilbkd") == 5 {
		t.Log("test ok")
	} else {
		t.Fail()
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LengthOfLongestSubstring("abcabcbb")
	}
}
