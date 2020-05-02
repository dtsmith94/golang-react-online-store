package main

import (
	"fmt"

	"github.com/dtsmith94/golang-react-online-store/models"
)

func main() {
	basket, _ := models.AddBasket(models.Basket{})

	fmt.Println("Basket ID:", basket.ID)

}
