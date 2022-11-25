package popularity

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Id string
type Title string
type Category string
type Popularity int

type PopularityRecord struct {
	Id         Id
	Title      Title
	Category   Category
	Popularity Popularity
}

type PopularitySource interface {
	Load(count int) ([]PopularityRecord, error)
}

type srgPopularitySource struct {
	client http.Client
	token  string
}

type srgSocialCountEntry struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type srgTopicEntry struct {
	Title string `json:"title"`
}

type srgShowEntry struct {
	Title     string          `json:"title"`
	TopicList []srgTopicEntry `json:"topicList"`
}

type srgMediaEntry struct {
	Id              string                `json:"id"`
	Title           string                `json:"title"`
	SocialCountList []srgSocialCountEntry `json:"socialCountList"`
	Show            srgShowEntry          `json:"show"`
}

type srgResponse struct {
	Next      string          `json:"next"`
	MediaList []srgMediaEntry `json:"mediaList"`
}

func (s srgPopularitySource) Load(count int) ([]PopularityRecord, error) {
	result := []PopularityRecord{}

	remaining := count
	base := "https://api.srgssr.ch/videometadata/v2/most_clicked?bu=srf"
	next := ""

	for remaining > 0 {
		pageUrl := base
		if next != "" {
			pageUrl = fmt.Sprintf("%s&next=%s", base, next)
		}
		req, reqErr := http.NewRequest("GET", pageUrl, nil)
		if reqErr != nil {
			return result, reqErr
		}

		req.Header.Add("Accept", "application/json")
		req.Header.Add("Authorization", strings.Join([]string{"Bearer", s.token}, " "))

		log.Println("Attempting to fetch", pageUrl)
		res, resErr := s.client.Do(req)
		if resErr != nil {
			return result, resErr
		}
		if res.StatusCode != 200 {
			return result, fmt.Errorf("failed to fetch resource: HTTP status=%d", res.StatusCode)
		}

		defer res.Body.Close()

		target := &srgResponse{}
		jsonErr := json.NewDecoder(res.Body).Decode(target)
		if jsonErr != nil {
			return result, jsonErr
		}

		for _, mediaEntry := range target.MediaList {
			category := ""
			popularity := -1

			for _, socialCountEntry := range mediaEntry.SocialCountList {
				if socialCountEntry.Key == "srgView" {
					popularity = socialCountEntry.Value
				}
			}

			for _, topicEntry := range mediaEntry.Show.TopicList {
				category = topicEntry.Title
			}

			rec := PopularityRecord{
				Id:         Id(mediaEntry.Id),
				Title:      Title(mediaEntry.Title),
				Category:   Category(category),
				Popularity: Popularity(popularity),
			}
			result = append(result, rec)
		}

		remaining -= len(target.MediaList)
		nextUrl, urlErr := url.Parse(target.Next)
		if urlErr != nil {
			return result, nil
		}
		next = nextUrl.Query().Get("next")
	}

	return result, nil
}

func NewPopularitySource() PopularitySource {
	client := http.Client{}
	token := os.Getenv("SRG_OAUTH_TOKEN")
	return srgPopularitySource{
		client,
		token,
	}
}
