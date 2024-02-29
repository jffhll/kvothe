package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var rarities *RarityRegistry = NewRarityRegistry()

func TestItemNew(t *testing.T) {
	assert := assert.New(t)
	i, error := NewItem("Short Sword", 10, "A common short sword, made of iron with a simple hilt", 5, false, rarities.Common)
	if error != nil {
		t.Fail()
	}
	check := &Item{
		name:        "Short Sword",
		price:       10,
		description: "A common short sword, made of iron with a simple hilt",
		stock:       5,
		magic:       false,
		rarity:      rarities.Common,
	}
	assert.Equal(*i, *check)
}
