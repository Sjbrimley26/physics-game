package arrayutils

// Sum returns the sum of an array of numbers
func Sum(arr []float64) float64 {
	total := 0.0
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	return total
}

// Min returns the smallest value of an array of numbers
func Min(arr []float64) float64 {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

// Max returns the greatest value of an array
func Max(arr []float64) float64 {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

type mmapper func(x float64) float64

// MMap returns a new array with all items changed by the mapping function.
func MMap(arr []float64, fn mmapper) []float64 {
	res := make([]float64, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = fn(arr[i])
	}
	return res
}
