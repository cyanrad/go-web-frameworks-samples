package coffee

import (
	"math/rand"
	"time"
)

func generateCoffeeData(count int) []Coffee {
	rand.Seed(time.Now().UnixNano())

	coffeeData := []Coffee{}

	for i := 1; i <= count; i++ {
		coffee := Coffee{
			ID:             i,
			InstantCoffee:  rand.Intn(10) + 1,
			CoffeeMate:     rand.Intn(5) + 1,
			PowderedMilk:   rand.Intn(5) + 1,
			EvaporatedMilk: rand.Intn(5) + 1,
			Water:          rand.Intn(100) + 1,
			Rating:         rand.Intn(5) + 1,
		}

		coffeeData = append(coffeeData, coffee)
	}

	return coffeeData
}
