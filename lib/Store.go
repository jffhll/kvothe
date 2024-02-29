package lib

type Store struct {
	name      string
	inventory []*Item
	currency  int
}

func NewStore(storeName string, storeInventory []*Item, storeCurrency int) (*Store, error) {
	store := &Store{
		name:      storeName,
		inventory: storeInventory,
		currency:  storeCurrency,
	}
	return store, nil
}

func (s *Store) AddItem(item *Item) error {
	s.inventory = append(s.inventory, item)
	return nil
}
