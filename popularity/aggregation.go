package popularity

import (
	"math"
	"sort"
)

type CategoryRanking map[Category]Score

func NewCountAggregation(records PopularityList) CategoryRanking {
	ranking := CategoryRanking{}
	for _, record := range records {
		ranking[record.Category] += 1
	}
	return ranking
}

func NewSumAggregation(records PopularityList) CategoryRanking {
	ranking := CategoryRanking{}
	for _, record := range records {
		ranking[record.Category] += Score(record.Popularity)
	}
	return ranking
}

func NewMaxAggregation(records PopularityList) CategoryRanking {
	ranking := CategoryRanking{}
	for _, record := range records {
		ranking[record.Category] = Score(math.Max(float64(ranking[record.Category]), float64(record.Popularity)))
	}
	return ranking
}

type OrderedCategoryEntry struct {
	Category
	Score
}

type OrderedCategoryList []OrderedCategoryEntry

func NewOrderedCategoryList(ranking CategoryRanking) OrderedCategoryList {
	result := OrderedCategoryList{}

	for category, score := range ranking {
		result = append(result, OrderedCategoryEntry{
			Category: category,
			Score:    score,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Score > result[j].Score
	})

	return result
}
