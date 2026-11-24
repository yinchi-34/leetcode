package solution

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func maxLength(nums []int) int {
	res := 2 // any subarray of length 2 satisfies prod == lcm * gcd
	prod := 1
	left := 0
	for right, x := range nums {
		// Shrink until x is coprime to every element in nums[left..right-1].
		for gcd(prod, x) > 1 {
			prod /= nums[left]
			left++
		}
		prod *= x
		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}
