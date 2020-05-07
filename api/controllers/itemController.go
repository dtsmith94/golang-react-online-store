package controllers

import (
	"net/http"

	"github.com/dtsmith94/golang-react-online-store/api/models"
)

// GetAllItems returns all the items available
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := models.GetAllItems()

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Items could not be returned due to an error")
		return
	}

	respondWithJSON(w, http.StatusOK, items)
}
