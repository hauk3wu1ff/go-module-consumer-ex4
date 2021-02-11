package main

import (
	"fmt"

	favs "github.com/hauk3wu1ff/go-module-ex4/quotes"
)

func favourites() []string {
	return favs.Favs()
}

func main() {
	fmt.Println("My Favourites are:")
	fmt.Println(favourites())
}
