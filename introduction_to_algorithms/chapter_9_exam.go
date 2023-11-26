package introduction_to_algorithms

import (
	"fmt"
	"math"
)

func run() {
	S := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Example set S
	k := 4                                        // Number of closest elements to find

	closestNumbers := findClosestNumbers(S, k)
	fmt.Println(closestNumbers)
}

func findClosestNumbers(S []float64, k int) []float64 {
	median := quickSelect2(S, 0, len(S)-1, len(S)/2) // Find the median using quickSelect2
	return selectClosest(S, median, k)               // Select k closest numbers to the median
}

func quickSelect2(arr []float64, left, right, k int) float64 {
	if left == right {
		return arr[left]
	}

	pivotIndex := partition(arr, left, right) // Partition the array around a pivot
	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return quickSelect2(arr, left, pivotIndex-1, k)
	} else {
		return quickSelect2(arr, pivotIndex+1, right, k)
	}
}

func partition(arr []float64, left, right int) int {
	pivot := arr[right]
	i := left

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]
	return i
}

func selectClosest(arr []float64, median float64, k int) []float64 {
	diff := make([]float64, len(arr))
	for i, num := range arr {
		diff[i] = math.Abs(num - median) // Calculate the absolute difference from the median
	}

	closest := make([]float64, k)
	copy(closest, arr[:k]) // Initialize the closest array with the first k elements of arr

	for i := k; i < len(arr); i++ {
		maxIndex := 0
		maxDiff := diff[0]

		// Find the index with the maximum difference in the closest array
		for j := 1; j < k; j++ {
			if diff[j] > maxDiff {
				maxIndex = j
				maxDiff = diff[j]
			}
		}

		// Replace the element with the maximum difference if the current element has a smaller difference
		if diff[i] < maxDiff {
			closest[maxIndex] = arr[i]
			diff[maxIndex] = diff[i]
		}
	}

	return closest
}
