package resolvers

import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/graphql-go/graphql"
)

func fetchArxivPaper(idList string) (model.Article, error) {
	apiURL := BuildAPIURL(idList, "", "", "", 0, 0)
	fmt.Println(apiURL)
	res, err := FetchResponse(apiURL)
	if err != nil {
		return model.Article{}, err
	}
	mv, err := mxj.NewMapXml(res)
	if err != nil {
		return model.Article{}, err
	}
	jsonRes, err := mv.Json()
	if err != nil {
		return model.Article{}, err
	}

	var response model.SingleResponse
	json.Unmarshal(jsonRes, &response)
	return response.SingleFeed.Entry, nil
}

var PaperResolver = func(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if !ok || strings.TrimSpace(id) == "" {
		id = ""
	}
	paper, err := fetchArxivPaper(id)
	if err != nil {
		return nil, err
	}

	return paper, nil
}
