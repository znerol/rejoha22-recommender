package main

import (
	"fmt"
	"log"

	"github.com/znerol/rejoha22-recommender/popularity"
)

func main() {
	source := popularity.NewPopularitySource()

	log.Println("Attempting to scrape 200 most popular videos")
	lines, err := source.Load(200)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Popularity\tCategory\tTitle\tId\n")
	for _, line := range lines {
		fmt.Printf("%d\t%s\t%s\t%s\n", line.Popularity, line.Category, line.Title, line.Id)
	}
}
