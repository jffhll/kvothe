package lib

type Shopkeeper struct {
	name        string
	description string
	haggleDC    int
	store       *Store
}

func NewShopkeeper(sk_name string, sk_desc string, sk_haggleDC int, sk_store *Store) (*Shopkeeper, error) {
	sk := &Shopkeeper{
		name:        sk_name,
		description: sk_desc,
		haggleDC:    sk_haggleDC,
		store:       sk_store,
	}
	return sk, nil
}

func (s *Shopkeeper) PurchaseItemByName(name string) {
	//TODO implement
}

func (s *Shopkeeper) BuyItem(item *Item) {
	//TODO implement
}
