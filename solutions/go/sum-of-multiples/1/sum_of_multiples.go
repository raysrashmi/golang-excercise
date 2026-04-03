package sumofmultiples


func SumMultiples(limit int, divisors ...int) int {
    seen := make(map[int]bool)

    for _, divisor := range divisors {
        if divisor == 0 {
            continue
        }
        for i := divisor; i < limit; i += divisor {
            seen[i] = true
        }
    }

    sum := 0
    for num := range seen {
        sum += num
    }

    return sum
}