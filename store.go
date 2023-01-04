package main

type storeItem struct {
	Name          string
	Price         int
	Description   string
	AlreadyBought bool
}

func fetchStore() []storeItem {
	return []storeItem{
		{
			Name:          "Fake Health Potion",
			Price:         4,
			Description:   "It pretends to heal you, but it actually doesn't ;).",
			AlreadyBought: false,
		},
		{
			Name:          "Crowbar",
			Price:         30,
			Description:   "I'm honestly not sure what you would use this for. There aren't any locks to pick, or even nails to pull out.",
			AlreadyBought: false,
		},
		{
			Name:          "Guitar",
			Price:         2,
			Description:   "It looks super duper cool. You won't be allowed to play it, though.",
			AlreadyBought: false,
		},
		{
			Name:          "Meat Pie",
			Price:         7,
			Description:   "It's a tasy pie, made out of meat. It doesn't do anything else.",
			AlreadyBought: false,
		},
		{
			Name:          "Penny Candy",
			Price:         1,
			Description:   "Candy that costs only a penny. Well... I guess it technically costs a Gold Piece.",
			AlreadyBought: false,
		},
	}
}