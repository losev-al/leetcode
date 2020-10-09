package main

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	left := 0
	right := 0
	test := make(map[rune]int)
	for _, rune := range t {
		test[rune] += 1
	}
	window := make(map[rune]int)
	input := []rune(s)
	answerLeft := 0
	answerRight := len(input)
	finded := 0
	for right < len(input) {
		window[input[right]] += 1
		if count, ok := test[input[right]]; ok && count == window[input[right]] {
			finded++
		}
		for finded == len(test) && left <= right {
			if right-left < answerRight-answerLeft {
				answerLeft = left
				answerRight = right
			}
			window[input[left]] -= 1
			if count, ok := test[input[left]]; ok && count == window[input[left]]+1 {
				finded--
			}
			left++
		}
		right++
		if right - left > answerRight - answerLeft {
			window[input[left]] -= 1
			if count, ok := test[input[left]]; ok && count == window[input[left]]+1 {
				finded--
			}
			left++
		}
	}
	if answerRight == len(input) {
		return ""
	}
	return string(input[answerLeft:answerRight+1])
}

//func minWindow(s string, t string) string {
//	if !exists(s, t) {
//		return ""
//	}
//	result := ""
//	left := 0
//	right := len(s)
//	for {
//		for exists(s[left:right], t) {
//			left++
//		}
//		left--
//		for exists(s[left:right], t) {
//			right--
//		}
//		if result == "" || len(result) > len(s[left:right+1]) {
//			result = s[left : right+1]
//		}
//		if left == 0 || left == right {
//			break
//		}
//		for !exists(s[left:right], t) {
//			left--
//			if left < 0 {
//				break
//			}
//		}
//		if left < 0 {
//			break
//		}
//	}
//	return result
//}
//
//func exists(s string, t string) bool {
//	for _, rune := range t {
//		index := strings.IndexRune(s, rune)
//		if index == -1 {
//			return false
//		} else {
//			s = s[:index] + s[index+1:]
//		}
//	}
//	return true
//}
