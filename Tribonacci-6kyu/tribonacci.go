package tribonacci

// NOTE:
// As the name may already reveal, it works basically like a Fibonacci, but summing the last 3 (instead of 2) numbers of the sequence to generate the next
// You need to create a fibonacci function that given a signature array/list, returns the first n elements - signature included of the so seeded sequence.

func Tribonacci(signature [3]float64, n int) []float64 {
	output := []float64{}

	if n > 0 {
		if n <= 3 {
			output = signature[:n]
		} else if n > 3 {
			output = signature[:3]
			sum := 0.0
			for _, ele := range output {
				sum += ele
			}
			for i := 3; i < n; i++ {
				output = append(output, sum)
				sum = 2*sum - output[i-3]
			}
		}
	}
	return output
}
