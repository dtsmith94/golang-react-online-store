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
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Keywords  []string `json:"keywords"`
	Price     int      `json:"price"`
	ShortDesc string   `json:"shortDesc"`
	LongDesc  string   `json:"longDesc"`
}

// GetItem returns an Item from the collection with the matching ID
func (i *Item) GetItem() error {
	for _, item := range items {
		if item.ID == i.ID {
			i = item
			return nil
		}
	}

	return fmt.Errorf("Item with ID '%v' not found", i.ID)
}

// CreateItem adds a new Item to the collection of Items
func (i *Item) CreateItem() error {
	if i.ID != 0 {
		return errors.New("New Item must not include ID or it must be set to zero")
	}

	i.ID = newItemID
	newItemID++
	items = append(items, i)
	return nil
}

// GetAllItems returns all items
func GetAllItems() ([]*Item, error) {
	return items, nil
}
