package leetcode

// https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/

func subarraysDivByK2(A []int, K int) int {
	var left, right, count int
	var maxRight = len(A)
	if maxRight == 0 {
		return 0
	}
	for left = 0; left < maxRight; left++ {
		for right = left; right < maxRight; right++ {
			sum := sumArr(A[left : right+1])
			if sum%K == 0 {
				count++
			}
		}
	}
	return count
}

func sumArr(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// 通常，涉及连续子数组问题的时候，我们使用前缀和来解决。
//
// 我们令 P[i] = A[0] + A[1] + ... + A[i]P[i]=A[0]+A[1]+...+A[i]。那么每个连续子数组的和 \textit{sum}(i, j)sum(i,j) 就可以写成 P[j] - P[i-1]P[j]−P[i−1]（其中 0 < i < j0<i<j）的形式。此时，判断子数组的和能否被 KK 整除就等价于判断 (P[j] - P[i-1]) \bmod K == 0(P[j]−P[i−1])modK==0，根据 同余定理，只要 P[j] \bmod K == P[i-1] \bmod KP[j]modK==P[i−1]modK，就可以保证上面的等式成立。
//
// 因此我们可以考虑对数组进行遍历，在遍历同时统计答案。当我们遍历到第 ii 个元素时，我们统计以 ii 结尾的符合条件的子数组个数。我们可以维护一个以前缀和模 KK 的值为键，出现次数为值的哈希表 \textit{record}record，在遍历的同时进行更新。这样在计算以 ii 结尾的符合条件的子数组个数时，根据上面的分析，答案即为 [0..i-1][0..i−1] 中前缀和模 KK 也为 P[i] \bmod KP[i]modK 的位置个数，即 \textit{record}[P[i] \bmod K]record[P[i]modK]。
//
// 最后的答案即为以每一个位置为数尾的符合条件的子数组个数之和。需要注意的一个边界条件是，我们需要对哈希表初始化，记录 \textit{record}[0] = 1record[0]=1，这样就考虑了前缀和本身被 KK 整除的情况。
//
// 注意：不同的语言负数取模的值不一定相同，有的语言为负数，对于这种情况需要特殊处理。
// 同余定理 https://baike.baidu.com/item/%E5%90%8C%E4%BD%99%E5%AE%9A%E7%90%86/1212360?fr=aladdin

func subarraysDivByK(A []int, K int) int {
	record := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, elem := range A {
		sum += elem
		modulus := (sum%K + K) % K
		ans += record[modulus]
		record[modulus]++
	}
	return ans
}
