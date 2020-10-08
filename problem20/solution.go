package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}

//func climbStairs(n int) int {
//	result := 0
//	i := 0
//	for i*2 <= n {
//		one := n - i*2
//		//step := int(fact(one +i)/fact(one)/fact(i))
//		//if step < 0 {
//		//	fmt.Println(i, one)
//		//}
//		//result += step
//		step := 1
//		for j := 1; j <= i; j++ {
//			step = step * (one + j) / j
//		}
//		result += step
//		i++
//	}
//	return result
//}

//var factCache map[int]uint = make(map[int]uint)
//
//func fact(n int) uint {
//	if result, ok := factCache[n]; ok {
//		return result
//	}
//	result := uint(1)
//	for i := 1; i <= n ; i++ {
//		result *= uint(i)
//		factCache[i] = result
//	}
//	return result
//}
