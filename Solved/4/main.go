package main

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var a float64
	var merged []int
	var i, j int
	for {
		if i >= len(nums1) && j >= len(nums2) {
			break
		} else if i >= len(nums1) && j <= len(nums2) {
			merged = append(merged, nums2[j])
		} else if i <= len(nums1) && j >= len(nums2) {
			merged = append(merged, nums1[i])
		} else if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			j--
		} else if nums1[i] > nums2[j] {
			merged = append(merged, nums2[j])
			i--
		} else if nums1[i] == nums2[j] {
			merged = append(merged, nums1[i])
			merged = append(merged, nums2[j])
		}
		j++
		i++
		// fmt.Println(merged, i, j)
	}
	if len(merged)%2 == 1 {
		a = float64(merged[(len(merged)-1)/2])
	} else {
		a = (float64(merged[(len(merged)-1)/2]) + float64(merged[(len(merged)+1)/2])) / 2
	}
	return a
}
