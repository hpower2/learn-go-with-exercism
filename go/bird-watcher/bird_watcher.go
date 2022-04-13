package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	temp := 0
	for _, v := range birdsPerDay{
		temp = temp + v
	}
	return temp
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	upBound := week*7
	lowBound := (week-1) * 7
	slices := birdsPerDay[lowBound:upBound]
	var value int
	for _, v := range slices{
		value = value + v
	}
	return value
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for in, _ := range birdsPerDay{
		if in % 2 == 0{
			birdsPerDay[in] = birdsPerDay[in] + 1
		}
		continue
	}
	return birdsPerDay
}
