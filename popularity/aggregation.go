package popularity

import "sort"

type CategoryRanking map[Category]Popularity

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
		ranking[record.Category] += record.Popularity
	}
	return ranking
}

type OrderedCategoryEntry struct {
	Popularity
	Category
}

type OrderedCategoryList []OrderedCategoryEntry

func NewOrderedCategoryList(ranking CategoryRanking) OrderedCategoryList {
	result := OrderedCategoryList{}

	for category, popularity := range ranking {
		result = append(result, OrderedCategoryEntry{
			Category:   category,
			Popularity: popularity,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Popularity > result[j].Popularity
	})

	return result
}
