package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://www.reddit.com/"
const subredditURL string = baseURL + "r/%s.json"
const subredditSearchURL string = baseURL + "r/%s/search.json"

type post struct {
	Data struct {
		URL      string `json:"url"`
		Stickied bool   `json:"stickied"`
	} `json:"data"`
}

type response struct {
	Body struct {
		Posts []post `json:"children"`
	} `json:"data"`
}

// TopSearch searches the given subreddit for the query and returns
// the top search result, or empty string if none was found.
func TopSearch(subreddit, query string) string {
	posts, err := search(subreddit, query)
	if err != nil {
		fmt.Println("Error making request: ", err)
		return ""
	}

	return topURLNotSticky(posts)
}

// Top returns the current top post of the given subreddit.
func Top(subreddit string) string {
	posts, err := top(subreddit)
	if err != nil {
		fmt.Println("Error making request: ", err)
		return ""
	}

	return topURLNotSticky(posts)
}

func topURLNotSticky(posts []post) string {
	for _, p := range posts {
		if !p.Data.Stickied {
			return p.Data.URL
		}
	}
	return ""
}

func search(subreddit, query string) ([]post, error) {
	u := fmt.Sprintf(subredditSearchURL, subreddit)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("q", query)
	q.Add("restrict_sr", "on")
	q.Add("sort", "hot")
	req.URL.RawQuery = q.Encode()
	return sendAndParseRequest(req)
}

func top(subreddit string) ([]post, error) {
	u := fmt.Sprintf(subredditURL, subreddit)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	return sendAndParseRequest(req)
}

func sendAndParseRequest(req *http.Request) ([]post, error) {
	req.Header.Set("User-Agent", "golang:dankbot:v0 (by /u/coolbrow)")
	client := http.Client{}
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
	return parse(b), nil
}

func parse(bytes []byte) []post {
	var result response
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Println("Error parsing json: ", err)
	}
	return result.Body.Posts
}
