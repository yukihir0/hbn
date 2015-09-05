package hbn

import "sort"

var empty struct{}

type Client struct {
	totalPages         int
	maxParallelRequest int
	api                api
}

func NewClient() Client {
	return Client{
		totalPages:         1,
		maxParallelRequest: 1,
		api:                NewHttpAPI(),
	}
}

func (c *Client) SetTotalPages(total int) {
	if total > 0 {
		c.totalPages = total
	}
}

func (c *Client) SetMaxParallelRequest(max int) {
	if max > 0 {
		c.maxParallelRequest = max
	}
}

func (c *Client) SetAPI(api api) {
	c.api = api
}

func (c Client) RequestBookmarks(user string) []Bookmark {
	bookmarksChan := make(chan []Bookmark)
	limitChan := make(chan struct{}, c.maxParallelRequest)

	go func() {
		for page := 0; page < c.totalPages; page++ {
			select {
			case limitChan <- empty:
				go func(user string, page int) {
					bookmarksChan <- c.api.RequestBookmarks(user, page)
					<-limitChan
				}(user, page)
			}
		}
	}()

	bookmarks := []Bookmark{}
	for i := 0; i < c.totalPages; i++ {
		bookmarks = append(bookmarks, <-bookmarksChan...)
	}

	return bookmarks
}

func (c Client) RequestFavoriteBookmarks(user string) []Bookmark {
	bookmarksChan := make(chan []Bookmark)
	limitChan := make(chan struct{}, c.maxParallelRequest)

	go func() {
		for page := 0; page < c.totalPages; page++ {
			select {
			case limitChan <- empty:
				go func(user string, page int) {
					bookmarksChan <- c.api.RequestFavoriteBookmarks(user, page)
					<-limitChan
				}(user, page)
			}
		}
	}()

	bookmarks := []Bookmark{}
	for i := 0; i < c.totalPages; i++ {
		bookmarks = append(bookmarks, <-bookmarksChan...)
	}

	return bookmarks
}

func (c Client) RequestHotEntryBookmarks() []Bookmark {
	return c.api.RequestHotEntryBookmarks()
}

func (c Client) RequestSearchBookmarks(query string) []Bookmark {
	return c.api.RequestSearchBookmarks(query)
}

func (c Client) RequestRelatedBookmarks(user string) []Bookmark {
	bookmarksChan := make(chan []Bookmark)
	limitChan := make(chan struct{}, c.maxParallelRequest)

	baseBookmarks := c.RequestBookmarks(user)
	go func() {
		for _, bookmark := range baseBookmarks {
			select {
			case limitChan <- empty:
				go func(url string) {
					bookmarksChan <- c.api.RequestRelatedBookmarks(url)
					<-limitChan
				}(bookmark.URL)
			}
		}
	}()

	bookmarks := []Bookmark{}
	for i := 0; i < len(baseBookmarks); i++ {
		bookmarks = append(bookmarks, <-bookmarksChan...)
	}

	return bookmarks
}

func (c Client) SearchNeighbors(bookmarks []Bookmark) Neighbors {
	ubChan := make(chan map[string]Bookmark)
	limitChan := make(chan struct{}, c.maxParallelRequest)

	go func() {
		for _, bookmark := range bookmarks {
			select {
			case limitChan <- empty:
				go func(url string) {
					ubChan <- c.api.RequestUserBookmark(url)
					<-limitChan
				}(bookmark.URL)
			}
		}
	}()

	// calculate common bookmarks
	cb := map[string][]Bookmark{}
	for i := 0; i < len(bookmarks); i++ {
		for u, b := range <-ubChan {
			cb[u] = append(cb[u], b)
		}
	}

	neighbors := Neighbors{}
	for user, common := range cb {
		neighbors = append(neighbors, Neighbor{
			User:            user,
			Similarity:      float64(len(common)) / float64(len(bookmarks)),
			CommonBookmarks: common,
		})
	}
	sort.Sort(neighbors)

	return neighbors
}
