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

	ranking := popularity.NewCountAggregation(lines)

	fmt.Printf("Popularity\tCategory\n")
	for _, entry := range popularity.NewOrderedCategoryList(ranking) {
		fmt.Printf("%.2f\t%s\n", entry.Score, entry.Category)
	}

}
