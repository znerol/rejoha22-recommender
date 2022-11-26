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

	fmt.Printf("Score\tPopularity\tCategory\tTitle\tId\n")
	for _, line := range popularity.NewNormalizedPopularityList(lines, 0, 1) {
		fmt.Printf("%f\t%f\t%s\t%s\t%s\n", line.Score, line.Popularity, line.Category, line.Title, line.Id)
	}
}
