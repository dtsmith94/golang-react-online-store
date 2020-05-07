package controllers

import (
	"net/http"
	"strconv"

	"github.com/dtsmith94/golang-react-online-store/api/models"
	"github.com/gorilla/mux"
)

// GetAllItemsInBasket gets all the items in a basket
func GetAllItemsInBasket(w http.ResponseWriter, r *http.Request) {
	routeParams := mux.Vars(r)
	basketID, err := strconv.Atoi(routeParams["basketId"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid basket ID")
	}

	basketItems, err := models.GetItemsInBasket(basketID)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Items not found for this basket")
		return
	}

	respondWithJSON(w, http.StatusOK, basketItems)
}

// AddItemToBasket adds an existing product to an existing basket
func AddItemToBasket(w http.ResponseWriter, r *http.Request) {
	routeParams := mux.Vars(r)
	basketID, err := strconv.Atoi(routeParams["basketId"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Basket ID")
	}

	itemID, err := strconv.Atoi(routeParams["itemId"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Item ID")
	}

	bi := models.BasketItem{
		BasketID: basketID,
		ItemID:   itemID,
	}

	err = bi.CreateBasketItem()

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Item or Basket ID does not exist")
		return
	}

	respondWithJSON(w, http.StatusOK, bi)
}
