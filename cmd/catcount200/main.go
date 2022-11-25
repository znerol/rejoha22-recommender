package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/znerol/rejoha22-recommender/popularity"
)

type Sum int

type rankEntry struct {
	Sum
	Category popularity.Category
}

func main() {
	source := popularity.NewPopularitySource()

	log.Println("Attempting to scrape 200 most popular videos")
	lines, err := source.Load(200)
	if err != nil {
		panic(err)
	}

	ranking := map[popularity.Category]Sum{}
	for _, line := range lines {
		ranking[line.Category] += 1
	}

	rankList := []rankEntry{}
	for category, sum := range ranking {
		rankList = append(rankList, rankEntry{
			Category: category,
			Sum:      sum,
		})
	}

	sort.Slice(rankList, func(i, j int) bool {
		return rankList[i].Sum > rankList[j].Sum
	})

	fmt.Printf("Sum\tCategory\n")
	for _, entry := range rankList {
		fmt.Printf("%d\t%s\n", entry.Sum, entry.Category)
	}

}
