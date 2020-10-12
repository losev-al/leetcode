package main

func addBinary(a string, b string) string {
	lenA := len(a)
	lenB := len(b)
	max := lenA
	if lenB > max {
		max = lenB
	}
	result := make([]byte, 0, max+1)
	overflow := 0
	for i := 0; i < max; i++ {
		first := 0
		if lenA - i > 0 && a[lenA - 1 - i] == '1' {
			first = 1
		}
		second := 0
		if lenB - i > 0 && b[lenB - 1 - i] == '1' {
			second = 1
		}
		result = append([]byte{byte(48 + (first+second+overflow)%2)}, result...)
		overflow = (first+second+overflow)/2
	}
	if overflow == 1 {
		result = append([]byte{byte(49)}, result...)
	}
	return string(result)
}
