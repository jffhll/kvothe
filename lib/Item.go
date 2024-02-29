package lib

type Item struct {
	name        string
	price       int
	description string
	stock       int
	magic       bool
	rarity      *Rarity
}

func NewItem(in_name string, in_price int, in_description string, in_stock int, in_magic bool, in_rarity *Rarity) (*Item, error) {
	item := &Item{
		name:        in_name,
		price:       in_price,
		description: in_description,
		stock:       in_stock,
		magic:       in_magic,
		rarity:      in_rarity,
	}
	return item, nil
}
