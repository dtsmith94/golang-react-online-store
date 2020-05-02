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
	ID int
}

// AddBasket adds a new basket to the collection of baskets
func AddBasket(basket Basket) (Basket, error) {
	if basket.ID != 0 {
		return Basket{}, errors.New("New Basket must not include ID or it must be set to zero")
	}

	basket.ID = newBasketID
	newBasketID++
	baskets = append(baskets, &basket)
	return basket, nil
}

// GetBasketByID returns a basket from the collection with the matching ID
func GetBasketByID(id int) (Basket, error) {
	for _, basket := range baskets {
		if basket.ID == id {
			return *basket, nil
		}
	}

	return Basket{}, fmt.Errorf("Basket with ID '%v' not found", id)
}
