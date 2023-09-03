package newsapi

import (
	"time"
)

type EverythingRequest struct {
	Q              string    `url:"q,omitempty"`
	QInTitle       []string  `url:"qInTitle,omitempty"`
	Sources        []string  `url:"sources,omitempty,comma"`
	Domains        []string  `url:"domains,omitempty,comma"`
	ExcludeDomains []string  `url:"excludeDomains,omitempty"`
	From           time.Time `url:"from,omitempty"`
	To             time.Time `url:"to,omitempty"`
	Language       string    `url:"language,omitempty"`
	SortBy         string    `url:"sortBy,omitempty"`
	PageSize       int       `url:"pageSize,omitempty"`
	Page           int       `url:"page,omitempty"`
}

type TopHeadlinesRequest struct {
	Country  string   `url:"country,omitempty"`
	Category string   `url:"category,omitempty"`
	Sources  []string `url:"sources,omitempty,comma"`
	Q        string   `url:"q, omitempty"`
	PageSize int      `url:"pageSize,omitempty"`
	Page     int      `url:"page,omitempty"`
}

type SourcesRequest struct {
	Category string `url:"category,omitempty"`
	Language string `url:"language,omitempty"`
	Country  string `url:"country,omitempty"`
}

type Response struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Sources     Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type SourcesResponse struct {
	Status  string   `json:"status"`
	Sources []Source `json:"sources"`
}

type Source struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}
