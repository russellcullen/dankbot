package haiku

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const search_haiku_url string = "https://www.reddit.com/r/youtubehaiku/search.json?q=%s&restrict_sr=on&sort=hot"

type post struct {
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
}

type response struct {
	Data struct {
		Children []post `json:"children"`
	} `json:"data"`
}

func TopUrl(query string) string {
	b, err := doRequest(query)
	if err != nil {
		fmt.Println("Error making request: ", err)
		return ""
	}

	var result response
	err = json.Unmarshal(b, &result)
	if err != nil {
		fmt.Println("Error parsing json: ", err)
		return ""
	}
	if len(result.Data.Children) == 0 {
		return ""
	}
	top := result.Data.Children[0]
	return top.Data.Url
}

func doRequest(query string) ([]byte, error) {
	client := http.Client{}

	newQuery := url.QueryEscape(query)
	req, err := http.NewRequest("GET", fmt.Sprintf(search_haiku_url, newQuery), nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	req.Header.Set("User-Agent", "golang:dankbot:v0 (by /u/coolbrow)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error searching for query: ", err)
		return nil, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response: ", err)
		return nil, err
	}
	return b, nil
}
