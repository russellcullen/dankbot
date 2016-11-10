package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
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

// RandomSearch searches the given subreddit for the query and returns
// a random result (from the first page), or empty string if none was found.
func RandomSearch(subreddit, query string) string {
	posts, err := search(subreddit, query)
	if err != nil {
		fmt.Println("Error making request: ", err)
		return ""
	}

	return randURLNotSticky(posts)
}

// Random returns a random post from the current first page of the given subreddit
func Random(subreddit string) string {
	posts, err := top(subreddit)
	if err != nil {
		fmt.Println("Error making request: ", err)
		return ""
	}

	return randURLNotSticky(posts)
}

func randURLNotSticky(posts []post) string {
	first := 0
	for _, p := range posts {
		if !p.Data.Stickied {
			break
		}
		first++
	}
	posts = posts[first:]
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	post := posts[r.Intn(len(posts))]
	return post.Data.URL
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
