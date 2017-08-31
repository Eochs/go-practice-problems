package main

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	mostAreaSeen := 0
	for left < right {
		hl := height[left]
		hr := height[right]
		if hl == hr {
			mostAreaSeen = max(hl*(right-left), mostAreaSeen)
			left++
			right--
		}
		if hl < hr {
			mostAreaSeen = max(hl*(right-left), mostAreaSeen)
			left++
		}
		if hr < hl {
			mostAreaSeen = max(hr*(right-left), mostAreaSeen)
			right--
		}
	}
	return mostAreaSeen
}

func main() {

}
