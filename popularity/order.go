package popularity

import (
	"math"
	"sort"
)

func NewNormalizedPopularityList(records PopularityList, expGeneral float64, expCategory float64) PopularityList {
	ranking := NewMaxAggregation(records)

	maxGeneralPopularity := 0.0
	for _, popularity := range ranking {
		maxGeneralPopularity = math.Max(maxGeneralPopularity, float64(popularity))
	}

	result := PopularityList{}
	for _, record := range records {
		maxCatPopularity := float64(ranking[record.Category])
		normalizedCatScore := float64(record.Popularity) / maxCatPopularity
		normalizedGeneralScore := float64(record.Popularity) / maxGeneralPopularity
		normalizedScore := math.Pow(normalizedGeneralScore, expGeneral) * math.Pow(normalizedCatScore, expCategory)
		result = append(result, PopularityRecord{
			Id:         record.Id,
			Title:      record.Title,
			Category:   record.Category,
			Popularity: record.Popularity,
			Score:      Score(normalizedScore),
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Score > result[j].Score
	})

	return result
}
