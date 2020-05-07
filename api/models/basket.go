package models

import (
	"errors"
	"fmt"
)

var (
	baskets     []*Basket
	newBasketID = 1
)

// Basket represents a collection of Items which can be ordered
type Basket struct {
	ID int `json:"id"`
}

// CreateBasket adds a new basket to the collection of baskets
func (b *Basket) CreateBasket() error {
	if b.ID != 0 {
		return errors.New("New Basket must not include ID or it must be set to zero")
	}

	b.ID = newBasketID
	newBasketID++
	baskets = append(baskets, b)
	return nil
}

// GetBasket returns a basket from the collection with the matching ID
func (b *Basket) GetBasket() error {
	for _, basket := range baskets {
		if basket.ID == b.ID {
			b = basket
			return nil
		}
	}

	return fmt.Errorf("Basket with ID '%v' not found", b.ID)
}
