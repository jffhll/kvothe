package lib

type Rarity struct {
	name      string
	magnitude int
	color     string //TODO switch away from stringly typing here
}

func NewRarity(in_name string, in_magnitude int, in_color string) *Rarity {
	return &Rarity{
		name:      in_name,
		magnitude: in_magnitude,
		color:     in_color,
	}
}
