package main

const (
	FIZZ     = 3
	BUZZ     = 5
	FIZZBUZZ = 15
)

func main() {

	for i := 1; i <= 100; i++ {
		switch {
		case (i % FIZZBUZZ) == 0:
			println("FizzBuzz")
		case (i % FIZZ) == 0:
			println("Fizz")
		case (i % BUZZ) == 0:
			println("Buzz")
		default:
			println(i)
		}
	}
}
