package birdwatcher

//import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    sum := 0
    for i:= 0; i < len(birdsPerDay) ;  i++ {
        sum = sum + birdsPerDay[i]
    }
    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    sum := 0
    startIndex := (week - 1) * 7
    endIndex := week * 7
     for i:= startIndex; i < endIndex ;  i++ {
        sum = sum + birdsPerDay[i]
    }
    return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i:= 0; i < len(birdsPerDay) ;  i++ {
        if i%2 == 0 || i==0{ 
            birdsPerDay[i] += 1
        }
        
    }
    return birdsPerDay
}
