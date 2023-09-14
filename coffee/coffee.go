package coffee

// all measurements are done in gram, hence int
type Coffee struct {
	ID             int
	InstantCoffee  int
	CoffeeMate     int
	PowderedMilk   int
	EvaporatedMilk int
	Water          int
	Rating         int
}

type CoffeeDB []Coffee

func (cdb CoffeeDB) init() {
	cdb = make([]Coffee, 0)
}

func (cdb CoffeeDB) Create(c Coffee) {
	newID := cdb[len(cdb)-1].ID + 1
	c.ID = newID // overwriting ID if for some reason written
	cdb = append(cdb, c)
}

func (cdb CoffeeDB) Get(ID int) (Coffee, bool) {
	if i, ok := cdb.findIndexFromID(ID); ok {
		return cdb[i], true
	}
	return Coffee{}, false
}

func (cdb CoffeeDB) Set(ID int, c Coffee) bool {
	if i, ok := cdb.findIndexFromID(ID); ok {
		cdb[i] = c
		return true
	}
	return false
}

func (cdb CoffeeDB) Delete(ID int) bool {
	if i, ok := cdb.findIndexFromID(ID); ok {
		cdb = append(cdb[:i], cdb[i+1:]...)
		return true
	}

	return false
}

func (cdb CoffeeDB) findIndexFromID(ID int) (int, bool) {
	for i, c := range cdb {
		if c.ID == ID {
			return i, true
		}
	}
	return -1, false
}

func (cdb CoffeeDB) Avg() Coffee {
	var avg Coffee

	for _, c := range cdb {
		avg.InstantCoffee += c.InstantCoffee
		avg.CoffeeMate += c.CoffeeMate
		avg.PowderedMilk += c.PowderedMilk
		avg.EvaporatedMilk += c.EvaporatedMilk
		avg.Water += c.Water
		avg.Rating += c.Rating
	}

	clen := len(cdb)
	avg.InstantCoffee /= clen
	avg.CoffeeMate /= clen
	avg.PowderedMilk /= clen
	avg.EvaporatedMilk /= clen
	avg.Water /= clen
	avg.Rating /= clen

	return avg
}
