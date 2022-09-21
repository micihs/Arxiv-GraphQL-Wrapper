package model

import "time"

type (
	
	Article struct {
		ID string
		Updated time.Time
		Published time.Time
		Title string
		Summary string
		Author []Author
		Link []Link
		PrimaryCatergory FeedCatergory
		Catergory []FeedCatergory
	}

	Response struct {
		Feed Feed
	}

	SingleResponse struct {
		SingleFeed singlefeed
	}

	Feed struct {
		Entry []Article
	}

	SingleFeed struct {
		Entry Article
	}

	Author struct {
		Name string
	}

	Link struct {
		Href string
		Rel string
		Title string
	}

	FeedCatergory struct {
		Term string
		Scheme string
	}

	Field struct {
		Name string
		Tag string
		Catergories []Catergory
	}

	Catergory struct {
		Name string
		Tag string
	}
)