package images

import (
	//"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	//"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const retroError string = "*Cannot generate retro image.*"

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Sombra returns an image of sombra dancing
func Sombra() string {
	return "http://i.imgur.com/lq3TwJi.gif"
}

// GenerateRIP generates a tombstone with the provided name
func GenerateRIP(name string) string {
	return fmt.Sprintf("http://www.tombstonebuilder.com/generate.php?top1=RIP&top3=%s", name)
}

// GenerateRetro generates a retro image with the text given.
func GenerateRetro(line1, line2, line3 string) string {
	form := url.Values{}
	form.Add("bcg", strconv.Itoa(random.Intn(5)+1))
	form.Add("txt", strconv.Itoa(random.Intn(4)+1))
	form.Add("text1", line1)
	form.Add("text2", line2)
	form.Add("text3", line3)
	resp, err := http.PostForm("https://photofunia.com/effects/retro-wave", form)
	if err != nil {
		fmt.Println("Error posting request: ", err)
		return retroError
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		fmt.Println("Error parsing retro result: ", err)
		return retroError
	}

	return doc.Find("#result-image").AttrOr("src", retroError)
}
