package hbn

import (
	"runtime"
	"sort"

	"github.com/yukihir0/hbapi"
)

type defaultStrategy struct {
	user               string
	totalPages         int
	maxParallelRequest int
}

func newDefaultStrategy(user string) defaultStrategy {
	s := defaultStrategy{
		totalPages:         1,
		maxParallelRequest: runtime.GOMAXPROCS(runtime.NumCPU()),
	}
	s.SetUser(user)
	return s
}

// SetUser set user.
func (s *defaultStrategy) SetUser(user string) {
	if user != "" {
		s.user = user
	}
}

// SetTotalPages set totalPages.
func (s *defaultStrategy) SetTotalPages(pages int) {
	if pages > 0 {
		s.totalPages = pages
	}
}

// SetMaxParallelRequest set maxParallelRequest.
func (s *defaultStrategy) SetMaxParallelRequest(max int) {
	if max > 0 {
		s.maxParallelRequest = max
	}
}

var empty struct{}

func (s defaultStrategy) getBookmarks() []string {
	bookmarksChan := make(chan []string)
	limitChan := make(chan struct{}, s.maxParallelRequest)

	go func() {
		for page := 0; page < s.totalPages; page++ {
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
				}(s.user, page)
			}
		}
	}()

	bookmarks := []string{}
	for i := 0; i < s.totalPages; i++ {
		bookmarks = append(bookmarks, <-bookmarksChan...)
	}

	return bookmarks
}

func (s defaultStrategy) calcNeighbors(bookmarks []string) Neighbors {
	infoChan := s.getEntryInfoChannel(bookmarks)

	all := []hbapi.Entry{}
	common := map[string][]hbapi.Entry{}
	for i := 0; i < len(bookmarks); i++ {
		info := <-infoChan
		all = append(all, info.Entry)
		for _, bookmark := range info.Bookmarks {
			common[bookmark.User] = append(common[bookmark.User], info.Entry)
		}
	}

	neighbors := Neighbors{}
	for k, v := range common {
		neighbors = append(neighbors, Neighbor{
			User:            k,
			CommonBookmarks: v,
			AllBookmarks:    all,
		})
	}
	sort.Sort(neighbors)

	return neighbors
}

func (s defaultStrategy) getEntryInfoChannel(urls []string) <-chan hbapi.EntryInfo {
	infoChan := make(chan hbapi.EntryInfo)
	limitChan := make(chan struct{}, s.maxParallelRequest)

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
