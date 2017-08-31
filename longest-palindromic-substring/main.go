package main

func longestPalindrome(s string) string {
	longest := s[:1]

	for og := range s { // outer loop, moving og index to be center of palindrome check
		i, j := og-1, og+1
		for { // loop until no more palindrome centered at og
			if i >= 0 && j < len(s) && s[i] == s[j] {
				// inner palindrome
				if len(s[i:j+1]) > len(longest) {
					longest = s[i : j+1]
				}
				i--
				j++
				continue
			}
			if i >= 0 && s[i] == s[og] {
				if len(s[i:og+1]) > len(longest) {
					longest = s[i : og+1]
				}
				i--
				continue
			}
			if j < len(s) && s[j] == s[og] {
				if len(s[og:j+1]) > len(longest) {
					longest = s[og : j+1]
				}
				j++
				continue
			}
			break // no more palindromes
		}

	}
	return longest
}

func main() {

}
