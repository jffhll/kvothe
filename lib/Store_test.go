package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUpTestStore() (*Store, error) {
	i1, _ := NewItem("Short Sword", 10, "A common short sword, made of iron with a simple hilt", 5, false, rarities.Common)
	i2, _ := NewItem("Long bow", 20, "A long bow with runes carved into the wood", 10, false, rarities.Common)
	i3, _ := NewItem("Locket of Lathander", 2000, "A magical locket in the shape of a sun; the symbol of the Dawnfather", 1, true, rarities.Legendary)

	new_inventory := []*Item{i1, i2, i3}
	return NewStore("The Prancing Pony", new_inventory, 1000)
}

func TestStoreNew(t *testing.T) {
	assert := assert.New(t)
	i1, _ := NewItem("Short Sword", 10, "A common short sword, made of iron with a simple hilt", 5, false, rarities.Common)
	i2, _ := NewItem("Long bow", 20, "A long bow with runes carved into the wood", 10, false, rarities.Common)
	i3, _ := NewItem("Locket of Lathander", 2000, "A magical locket in the shape of a sun; the symbol of the Dawnfather", 1, true, rarities.Legendary)

	new_inventory := []*Item{i1, i2, i3}
	s, error := NewStore("The Prancing Pony", new_inventory, 1000)
	if error != nil {
		t.Fail()
	}
	check := &Store{
		name:      "The Prancing Pony",
		inventory: new_inventory,
		currency:  1000,
	}
	assert.Equal(*s, *check)
}

func TestAddItem(t *testing.T) {
	s, err := setUpTestStore()
	if err != nil {
		t.Fail()
	}
	new, _ := NewItem("Frostmourne", 4000, "A cursed runeblade bearing the evil soul of Ner'zhul", 1, true, rarities.Legendary)
	s.AddItem(new)
	updatedInventory, _ := s.getInventory()
	found := false
	for _, item := range updatedInventory {
		if item == new {
			found = true
		}
	}
	if !found {
		t.Fail()
	}
}
