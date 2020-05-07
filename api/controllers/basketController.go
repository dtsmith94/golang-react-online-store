package controllers

import (
	"net/http"

	"github.com/dtsmith94/golang-react-online-store/api/models"
)

// CreateBasket handles a request to create a new Basket
func CreateBasket(w http.ResponseWriter, r *http.Request) {
	b := models.Basket{}
	err := b.CreateBasket()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to create new Basket")
		return
	}

	respondWithJSON(w, http.StatusCreated, b)
}
