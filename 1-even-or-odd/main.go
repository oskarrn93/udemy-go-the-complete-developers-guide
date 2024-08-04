package main

const ARRAY_LENGTH = 11

func isEven(x int) bool {
	return x%2 == 0
}

func main() {
	values := make([]int, ARRAY_LENGTH)

	// populate array with values
	for i := range values {
		values[i] = i
	}

	for _, value := range values {
		print(value, ": ")

		if isEven(value) {
			println("even")
		} else {
			println("odd")
		}
	}

}
