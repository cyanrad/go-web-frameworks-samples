package main

// all measurements are done in gram, hence int
type Coffee struct {
	InstantCoffee  int
	CoffeeMate     int
	PowderedMilk   int
	EvaporatedMilk int
	Water          int
	Rating         int
}

// global variable to contain coffee data
// cuz i don't feel like handling databases
//
// if you think about, if you just keep the program running forever
// you won't ever need a database again :))))))
var coffees []Coffee = make([]Coffee, 0)

func CoffeeAvg() (avg Coffee) {
	for _, c := range coffees {
		avg.InstantCoffee += c.InstantCoffee
		avg.CoffeeMate += c.CoffeeMate
		avg.PowderedMilk += c.PowderedMilk
		avg.EvaporatedMilk += c.EvaporatedMilk
		avg.Water += c.Water
		avg.Rating += c.Rating
	}

	clen := len(coffees)
	avg.InstantCoffee /= clen
	avg.CoffeeMate /= clen
	avg.PowderedMilk /= clen
	avg.EvaporatedMilk /= clen
	avg.Water /= clen
	avg.Rating /= clen

	return avg
}
