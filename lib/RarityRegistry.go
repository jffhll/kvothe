package lib

type RarityRegistry struct {
	rarities []*Rarity
}

func NewRarityRegistry() *RarityRegistry {
	// Set up accepted item rarites
	common := NewRarity("Common", 1, "grey")
	uncommon := NewRarity("Uncommon", 2, "green")
	rare := NewRarity("Rare", 3, "blue")
	veryRare := NewRarity("Very Rare", 4, "purple")
	legendary := NewRarity("Legendary", 5, "yellow")

	return &RarityRegistry{
		rarities: []*Rarity{
			common,
			uncommon,
			rare,
			veryRare,
			legendary,
		},
	}
}
