package popularity

import "sort"

func NewNormalizedPopularityList(records PopularityList) PopularityList {
	ranking := NewMaxAggregation(records)

	result := PopularityList{}
	for _, record := range records {
		maxScore := ranking[record.Category]
		normalizedScore := Score(record.Popularity) / maxScore
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
