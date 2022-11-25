package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/znerol/rejoha22-recommender/popularity"
)

type rankEntry struct {
	Popularity popularity.Popularity
	Category   popularity.Category
}

func main() {
	source := popularity.NewPopularitySource()

	log.Println("Attempting to scrape 200 most popular videos")
	lines, err := source.Load(200)
	if err != nil {
		panic(err)
	}

	ranking := map[popularity.Category]popularity.Popularity{}
	for _, line := range lines {
		ranking[line.Category] += line.Popularity
	}

	rankList := []rankEntry{}
	for category, popularity := range ranking {
		rankList = append(rankList, rankEntry{
			Category:   category,
			Popularity: popularity,
		})
	}

	sort.Slice(rankList, func(i, j int) bool {
		return rankList[i].Popularity > rankList[j].Popularity
	})

	fmt.Printf("Popularity\tCategory\n")
	for _, entry := range rankList {
		fmt.Printf("%d\t%s\n", entry.Popularity, entry.Category)
	}

}
