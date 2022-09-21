package resolvers

import (
	"fmt"
	"strings"
	"encoding/json"
	"github.com/graphql-go/graphwl"
)

const (
	SortOrderAscending = "ascending"
	SortOrderDescending = "descending"
	SortByRelevance = "relevance"
	SortByLastUpdatedDate = "lastUpdateDate"
	SortBySubmittedDate = "submittedDate"
	resultLimit = 10
)

func fetchArxivFeed(catergory, sortBY, sortOrder string, limit, offset int) ([]model.Article, error) {
	APIURL := BuidlAPIURL("", catergory, sortBy, sortOrder, offset, limit)
	fmt.Println(APIURL)
	res, err := FetchResponse(APIURL)
	if err != nil {
		return nil, err
	}

	mv, err := nxj.NewMapXml(res)
	if err != nil {
		return nil, err
	}

	jsonRes, err := mv.Json()
	if err != nil {
		return nil, err
	}

	var response model.Response
	json.Unmarshall(jsonRes, &response)
	return response.Feed.Entry, nil
}

var FeedResolver = func(p graphql) (interface{}, error) {
	limit, ok := p.Args["limit"].(int)
	if !ok || limit > resultsLimit || limit < 1 {
		limit = resultsLimit
	}
	offset, ok := p.Args["offset"].(int)
	if !ok || offset < 1 {
		offset = 0
	}
	sortBy, ok := p.Args["sortBy"].(string)
	if !ok || strings.TrimSpace(sortBy) == "" {
		sortBy = SortBySubmittedDate
	}
	sortOrder, ok := p.Args["sortOrder"].(string)
	if !ok || strings.TrimSpace(sortOrder) == "" {
		sortOrder = SortOrderDescending
	}
	category, ok := p.Args["category"].(string)
	if !ok || strings.TrimSpace(category) == "" {
		category = "cs.AI"
	}

	feedEntries, err := fetchArxivFeed(category, sortBy, sortOrder, limit, offset)
	if err != nil {
		return nil, err
	}

	return feedEntries, nil
}