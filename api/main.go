package main

import (
	"log"
	"net/http"

	"github.com/dtsmith94/golang-react-online-store/api/controllers"
	"github.com/dtsmith94/golang-react-online-store/api/models"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	initializeRoutes(router)

	mockItems(10)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func initializeRoutes(router *mux.Router) {
	router.HandleFunc("/baskets/{basketId:[0-9]+}/items", controllers.GetAllItemsInBasket).Methods("GET")
	router.HandleFunc("/baskets/{basketId:[0-9]+}/items/{itemId:[0-9]+}", controllers.AddItemToBasket).Methods("PUT")
	router.HandleFunc("/baskets", controllers.CreateBasket).Methods("POST")
	router.HandleFunc("/items", controllers.GetAllItems).Methods("GET")
}

func mockItems(quantity int) {

	for i := 0; i < quantity; i++ {
		item := models.Item{
			Name:      "Item 1",
			Keywords:  []string{"cool", "exciting", "fun"},
			Price:     (i + 1) * 1000,
			ShortDesc: "Lorem Ipsum is simply dummy text",
			LongDesc:  "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
		}

		item.CreateItem()
	}

}
