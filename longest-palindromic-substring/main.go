package main

func expandAroundCenter(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func longestPalindrome(s string) string {
	longest := s[:1] //grab first character by default
	for ix := range s[:len(s)-1] {
		substr1 := expandAroundCenter(s, ix, ix)
		substr2 := expandAroundCenter(s, ix, ix+1)
		if len(substr1) > len(longest) {
			longest = substr1
		}
		if len(substr2) > len(longest) {
			longest = substr2
		}
	}
	return longest
}

// func longestPalindrome(s string) string {
// 	longest := s[:1]

// 	for og := range s { // outer loop, moving og index to be center of palindrome check
// 		i, j := og-1, og+1
// 		for { // loop until no more palindrome centered at og
// 			if i >= 0 && j < len(s) && s[i] == s[j] {
// 				// inner palindrome
// 				if len(s[i:j+1]) > len(longest) {
// 					longest = s[i : j+1]
// 				}
// 				i--
// 				j++
// 				continue
// 			}
// 			if i >= 0 && s[i] == s[og] {
// 				if len(s[i:og+1]) > len(longest) {
// 					longest = s[i : og+1]
// 				}
// 				i--
// 				continue
// 			}
// 			if j < len(s) && s[j] == s[og] {
// 				if len(s[og:j+1]) > len(longest) {
// 					longest = s[og : j+1]
// 				}
// 				j++
// 				continue
// 			}
// 			break // no more palindromes
// 		}

// 	}
// 	return longest
// }

func main() {

}
