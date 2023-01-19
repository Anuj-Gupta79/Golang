package main

func SumMultiples(limit int, divisors ...int) (sum int) {
	for i := 1; i < limit; i++ {
		for _, div := range divisors {
			if div > 0 && i%div == 0 {
				sum += i
				break
			}
		}
	}
	return
}

func main() {

}
