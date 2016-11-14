package textapis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CatFact returns a random cat fact
func CatFact() string {
	type response struct {
		Facts []string `json:"facts"`
	}

	var r response
	err := getAndParse("http://catfacts-api.appspot.com/api/facts", &r)
	if err != nil {
		return "*Error getting cat facts.*"
	}

	return r.Facts[0]
}

// Swanson returns a random Ron Swanson quote
func Swanson() string {
	var r []string
	err := getAndParse("http://ron-swanson-quotes.herokuapp.com/v2/quotes", &r)
	if err != nil {
		return "*Error gettin Swanson quote.*"
	}

	return r[0]
}

func getAndParse(u string, t interface{}) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &t)
	if err != nil {
		return err
	}

	return nil
}
