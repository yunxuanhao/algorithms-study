package leetcode

// https://leetcode-cn.com/problems/find-in-mountain-array/

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

type MountainArray struct {
	arr []int
}

func (m *MountainArray) get(index int) int {
	if index < 0 || index >= len(m.arr) {
		return -1
	}
	return m.arr[index]
}

func (m *MountainArray) length() int {
	return len(m.arr)
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	l, r := 0, mountainArr.length()-1
	for l < r {
		mid := (l + r) / 2
		midNum := mountainArr.get(mid)
		midAddOneNum := mountainArr.get(mid + 1)
		if midNum == target {
			return mid
		}

		if midAddOneNum-midNum > 0 {
			if midNum > target {
				r = mid
			} else {
				l = mid + 1
			}
		} else {
			if midNum > target {
				r = mid
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
