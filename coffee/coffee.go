package coffee

// all measurements are done in gram, hence int
type Coffee struct {
	ID             int `json:"id"`
	InstantCoffee  int `json:"instant_coffee"`
	CoffeeMate     int `json:"coffee_mate"`
	PowderedMilk   int `json:"powdered_milk"`
	EvaporatedMilk int `json:"evaporated_milk"`
	Water          int `json:"water"`
	Rating         int `json:"rating"`
}

type CoffeePatch struct {
	InstantCoffee  *int `json:"instant_coffee"`
	CoffeeMate     *int `json:"coffee_mate"`
	PowderedMilk   *int `json:"powdered_milk"`
	EvaporatedMilk *int `json:"evaporated_milk"`
	Water          *int `json:"water"`
	Rating         *int `json:"rating"`
}

type CoffeeDB []Coffee

func (cdb *CoffeeDB) Init() {
	*cdb = generateCoffeeData(10)
}

func (cdb *CoffeeDB) Create(c Coffee) {
	newID := (*cdb)[len(*cdb)-1].ID + 1
	c.ID = newID // overwriting ID if for some reason written
	*cdb = append(*cdb, c)
}

func (cdb CoffeeDB) Get(ID int) (Coffee, bool) {
	if i, ok := cdb.findIndexFromID(ID); ok {
		return cdb[i], true
	}
	return Coffee{}, false
}

func (cdb *CoffeeDB) Delete(ID int) bool {
	if i, ok := cdb.findIndexFromID(ID); ok {
		*cdb = append((*cdb)[:i], (*cdb)[i+1:]...)
		return true
	}

	return false
}

func (cdb *CoffeeDB) Patch(ID int, patch CoffeePatch) (Coffee, bool) {
	i, ok := cdb.findIndexFromID(ID)
	if !ok {
		return Coffee{}, false
	}

	coffee := (*cdb)[i]

	if patch.InstantCoffee != nil {
		coffee.InstantCoffee = *patch.InstantCoffee
	}
	if patch.CoffeeMate != nil {
		coffee.CoffeeMate = *patch.CoffeeMate
	}
	if patch.PowderedMilk != nil {
		coffee.PowderedMilk = *patch.PowderedMilk
	}
	if patch.EvaporatedMilk != nil {
		coffee.EvaporatedMilk = *patch.EvaporatedMilk
	}
	if patch.Water != nil {
		coffee.Water = *patch.Water
	}
	if patch.Rating != nil {
		coffee.Rating = *patch.Rating
	}

	(*cdb)[i] = coffee
	return coffee, true
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
