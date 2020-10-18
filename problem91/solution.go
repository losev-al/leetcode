package main

import "strconv"

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	prev := 1
	current := 1
	next := 1
	for i := 1; i < len(s); i++ {
		num, _ := strconv.ParseInt(s[i-1:i+1], 10, 32)
		if s[i] == '0' && num != 10 && num != 20 {
			return 0
		} else if s[i] == '0' {
			next = prev
		} else if num < 11 || num > 26 {
			next = current
		} else {
			next = current + prev
		}
		prev = current
		current = next
	}
	return next
}

//func numDecodings(s string) int {
//	if strings.Index(s, "00") > 0 || s[0] == '0' {
//		return 0
//	}
//	count := 0
//	flag := true
//	calc(s, &count, &flag)
//	if flag {
//		return count
//	} else {
//		return 0
//	}
//}
//
//func calc(s string, count *int, flag *bool) {
//	if s == "" {
//		*count += 1
//	} else if s[0] == '0' {
//	} else if *flag {
//		calc(s[1:], count, flag)
//		if len(s) > 1 && (s[0] == '1' || s[0] == '2' && s[1] < '7') {
//			calc(s[2:], count, flag)
//		}
//	}
//}
