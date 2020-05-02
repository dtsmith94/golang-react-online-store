package models

import "errors"

var (
	basketItems []*BasketItem
)

// BasketItem represents the relationship between Baskets and Items
type BasketItem struct {
	BasketID int
	ItemID   int
	Quantity int
}

// AddBasketItem associates an item with a basket
func AddBasketItem(basketID int, itemID int) (BasketItem, error) {

	// check the basket and item exist first
	_, error := GetBasketByID(basketID)

	if error != nil {
		return BasketItem{}, error
	}

	_, error = GetItemByID(itemID)

	if error != nil {
		return BasketItem{}, error
	}

	// check if this item has already been added to this basket
	newBasketItem, error := getBasketItem(basketID, itemID)

	// if a BasketItem was found, increase quantity by one
	if error == nil {
		newBasketItem.Quantity++
	} else {
		newBasketItem = BasketItem{
			BasketID: basketID,
			ItemID:   itemID,
			Quantity: 1,
		}

		basketItems = append(basketItems, &newBasketItem)
	}

	return newBasketItem, nil
}

func getBasketItem(basketID int, itemID int) (BasketItem, error) {
	for _, basketItem := range basketItems {
		if basketItem.BasketID == basketID && basketItem.ItemID == itemID {
			return *basketItem, nil
		}
	}

	return BasketItem{}, errors.New("Unable to find BasketItem")
}
