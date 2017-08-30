package main

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

//LengthOfLongestSubstring ...
func LengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	hashmap := map[byte]int{}
	i := 0
	for j := 0; j < n; j++ {
		if v, ok := hashmap[s[j]]; ok { // s[j] is a duplicate value
			i = max(v, i)
		}
		ans = max(ans, j-i+1)
		hashmap[s[j]] = j + 1
	}
	return ans
}

func main() {

}
