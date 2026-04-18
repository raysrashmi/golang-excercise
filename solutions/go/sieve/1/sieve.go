package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	isPrime := make([]bool, limit+1)

	// assume all >=2 are prime
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	// sieve
	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}

	var primes []int

	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}