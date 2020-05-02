package models

import (
	"errors"
	"fmt"
)

var (
	items     []*Item
	newItemID = 1
)

// Item represents a product which can be browsed and added to a customer's basket
type Item struct {
	ID                  int
	Name                string
	Keywords            []string
	Price               int
	StockLevel          int
	IsActive            bool
	ShortDesc, LongDesc string
}

// GetItemByID returns an Item from the collection with the matching ID
func GetItemByID(id int) (Item, error) {
	for _, item := range items {
		if item.ID == id {
			return *item, nil
		}
	}

	return Item{}, fmt.Errorf("Item with ID '%v' not found", id)
}

// AddItem adds a new Item to the collection of Items
func AddItem(item Item) (Item, error) {
	if item.ID != 0 {
		return Item{}, errors.New("New Item must not include ID or it must be set to zero")
	}

	item.ID = newItemID
	newItemID++
	items = append(items, &item)
	return item, nil
}
