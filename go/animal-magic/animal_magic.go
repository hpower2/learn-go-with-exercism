package chance

import (
	"math/rand"
	"time"
)

var r int

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
	var die = make([]int, 20)
	for i := 0; i < len(die); i++ {
		die[i] = i + 1
	}
	return die[rand.Intn(len(die))]
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	return rand.Float64() * 12
}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
    animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
    for i := range animals {
        j := rand.Intn(i + 1)
        animals[i], animals[j] = animals[j], animals[i]
    }
    return animals
}