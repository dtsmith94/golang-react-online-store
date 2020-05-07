package models

var (
	basketItems []*BasketItem
)

// BasketItem represents the relationship between Baskets and Items
type BasketItem struct {
	Item     Item `json:"item"`
	BasketID int  `json:"basketId"`
	ItemID   int  `json:"itemId"`
	Quantity int  `json:"quantity"`
}

// CreateBasketItem associates an item with a basket
func (bi *BasketItem) CreateBasketItem() error {

	// check the basket and item exist first
	basket := Basket{ID: bi.BasketID}
	err := basket.GetBasket()

	if err != nil {
		return err
	}

	item := Item{ID: bi.ItemID}
	err = item.GetItem()

	if err != nil {
		return err
	}

	// check if this item has already been added to this basket. If it has increment quantity by 1
	itemAlreadyInBasket := false
	for i, basketItem := range basketItems {
		if basketItem.BasketID == bi.BasketID && basketItem.ItemID == bi.ItemID {
			itemAlreadyInBasket = true

			basketItems[i].Quantity++
			bi.Quantity = basketItems[i].Quantity
			return nil
		}
	}

	// if item not already in basket, set quanitity to 1 and add to list
	if !itemAlreadyInBasket {
		bi.Quantity = 1
		basketItems = append(basketItems, bi)
	}

	return nil
}

// GetItemsInBasket gets all the items which are in a specific basket
func GetItemsInBasket(basketID int) ([]BasketItem, error) {

	// check the basket exists first
	basket := Basket{ID: basketID}
	err := basket.GetBasket()

	if err != nil {
		return []BasketItem{}, err
	}

	filteredBasketItems := []BasketItem{}

	for _, basketItem := range basketItems {
		if basketItem.BasketID == basketID {

			item := Item{ID: basketItem.ItemID}
			err := item.GetItem()

			if err != nil {
				return []BasketItem{}, err
			}

			basketItem.Item = item

			filteredBasketItems = append(filteredBasketItems, *basketItem)
		}
	}

	return filteredBasketItems, nil
}
