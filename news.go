package newsapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	topHeadline = "/v2/top-headlines"
)

type News struct {
	apiKey     string
	keyType    string
	scheme     string
	host       string
	httpClient *http.Client
}

func NewsAPI(apiKey string) *News {
	news := &News{
		apiKey:  apiKey,
		keyType: "X-Api-Key",
		scheme:  "https",
		host:    "newsapi.org",
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
	return news
}

func (news *News) urlBuild(endpoint string) *url.URL {
	url := &url.URL{
		Scheme: news.scheme,
		Host:   news.host,
		Path:   endpoint,
	}
	q := url.Query()
	q.Add("country", "us")
	url.RawQuery = q.Encode()
	return url
}

func (news *News) buildRequest(url *url.URL) *http.Request {
	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set(news.keyType, news.apiKey)
	return request
}

func (news *News) getRequest(request *http.Request) *http.Response {
	response, err := news.httpClient.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	return response
}

func (news *News) GetTop() {
	url := news.urlBuild(topHeadline)
	request := news.buildRequest(url)
	response := news.getRequest(request)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var r Response
	json.Unmarshal(body, &r)
	for i, article := range r.Articles {
		if article.Title != "[Removed]" && i < 10 {
			fmt.Printf("\n%d) %s\n\t%s0f\n", i+1, r.Articles[i].Title, r.Articles[i].URL)
		}
	}
	fmt.Println()
}
