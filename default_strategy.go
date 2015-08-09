package hbn

import (
	"runtime"
	"sort"

	"github.com/yukihir0/hbapi"
)

type defaultStrategy struct {
	User               string
	TotalPages         int
	MaxParallelRequest int
}

func newDefaultStrategy(user string) defaultStrategy {
	return defaultStrategy{
		User:               user,
		TotalPages:         1,
		MaxParallelRequest: runtime.GOMAXPROCS(runtime.NumCPU()),
	}
}

var empty struct{}

func (s defaultStrategy) getBookmarks() []string {
	bookmarksChan := make(chan []string)
	limitChan := make(chan struct{}, s.MaxParallelRequest)

	go func() {
		for page := 0; page < s.TotalPages; page++ {
			select {
			case limitChan <- empty:
				go func(user string, page int) {
					params := hbapi.NewFeedParams(user)
					params.SetPage(page)
					feed, err := hbapi.GetFeed(params)
					if err != nil {
						bookmarksChan <- []string{}
						<-limitChan
						return
					}

					urls := []string{}
					for _, item := range feed.Items {
						urls = append(urls, item.Link)
					}
					bookmarksChan <- urls
					<-limitChan
				}(s.User, page)
			}
		}
	}()

	bookmarks := []string{}
	for i := 0; i < s.TotalPages; i++ {
		bookmarks = append(bookmarks, <-bookmarksChan...)
	}

	return bookmarks
}

func (s defaultStrategy) getRelatedBookmarks() []string {
	bookmarks := s.getBookmarks()
	infoChan := s.getEntryInfoChannel(bookmarks)

	relatedBookmarks := []string{}
	for i := 0; i < len(bookmarks); i++ {
		entry := <-infoChan
		for _, related := range entry.Related {
			relatedBookmarks = append(relatedBookmarks, related.URL)
		}
	}

	return relatedBookmarks
}

func (s defaultStrategy) getHotEntryBookmarks() []string {
	bookmarks := []string{}

	params := hbapi.NewHotEntryFeedParams()
	feed, err := hbapi.GetHotEntryFeed(params)
	if err != nil {
		return bookmarks
	}

	for _, item := range feed.Items {
		bookmarks = append(bookmarks, item.Link)
	}

	return bookmarks
}

func (s defaultStrategy) getEntryInfoChannel(urls []string) <-chan hbapi.EntryInfo {
	infoChan := make(chan hbapi.EntryInfo)
	limitChan := make(chan struct{}, s.MaxParallelRequest)

	go func() {
		for _, url := range urls {
			select {
			case limitChan <- empty:
				go func(url string) {
					entry, err := hbapi.GetEntryInfo(url)
					if err != nil {
						infoChan <- hbapi.EntryInfo{}
						<-limitChan
						return
					}

					infoChan <- entry
					<-limitChan
				}(url)
			}
		}
	}()

	return infoChan
}

func (s defaultStrategy) calcNeighbors(bookmarks []string) Neighbors {
	infoChan := s.getEntryInfoChannel(bookmarks)

	total := []hbapi.Entry{}
	common := map[string][]hbapi.Entry{}
	for i := 0; i < len(bookmarks); i++ {
		info := <-infoChan
		total = append(total, info.Entry)
		for _, bookmark := range info.Bookmarks {
			common[bookmark.User] = append(common[bookmark.User], info.Entry)
		}
	}

	neighbors := Neighbors{}
	for k, v := range common {
		neighbors = append(neighbors, Neighbor{
			User:            k,
			CommonBookmarks: v,
			TotalBookmarks:  total,
		})
	}
	sort.Sort(neighbors)

	return neighbors
}
