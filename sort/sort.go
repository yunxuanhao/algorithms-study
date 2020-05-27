package sort

// mergeSort 归并排序
func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := 0 + (len(nums))/2

	list1 := mergeSort(nums[:mid])
	list2 := mergeSort(nums[mid:])

	var tmp []int
	var i, j int
	for i < len(list1) && j < len(list2) {
		if list1[i] > list2[j] {
			tmp = append(tmp, list1[i])
			i++
		} else {
			tmp = append(tmp, list2[j])
			j++
		}
	}

	for i < len(list1) {
		tmp = append(tmp, list1[i])
		i++
	}

	for j < len(list2) {
		tmp = append(tmp, list2[j])
		j++
	}

	return tmp
}

// heapSort 堆排序
func heapSort(nums []int) []int {
	last := len(nums) - 1
	for last >= 0 {
		for i := last; i > 0; i-- {
			if nums[i] > nums[(i-1)/2] {
				nums[i], nums[(i-1)/2] = nums[(i-1)/2], nums[i]
			}
		}
		nums[last], nums[0] = nums[0], nums[last]
		last--
	}
	return nums
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {
			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
