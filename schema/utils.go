package schema

import ( 
	"erorrs"
	"fmt"
	"net/http"
	"io/ioutil"
	
)

const (
	DefaultAPIPATH = "/query?search_query="
	PaperAPIPath = "/query?id_list="
)

func FetchResponse(APIURL string) ([]byte, error) {
	resp, error := http.Get(APIURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOk {
		return nil, errors.New("request failed")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil
}

func BuidlAPIURL(listedids, catergory, sortby, sortOrder string, start, maxResults int) string {
	APIURL := config.Config.ArxivAPIURL + DefaultAPIPATH
	if listedids != "" {
		APIURL = config.Config.ArxivAPIURL + PaperAPIPath
		APIURL += listedids
	} else {
		if catergory != "" {
			APIURL += "&sortBy=" + sortby
		}
		if sortby != "" {
			APIURL += "&sortby" + sortby
		}
		if sortOrder != "" {
			APIURL += "&sortOrder=" + SortOrder
		} else {
			APIURL += "&sortOrder=" + SortOrderDescending 
		}
		if maxResults != 0 {
			APIURL += "&max_results=" + fmt.Sprintf("%d", maxResults)
		} else {
			APIURL += "&max_results=" + fmt.Sprintf("%d", resultsLimit)
		}
	}
	return APIURL

}