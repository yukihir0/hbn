package hbn

import "github.com/yukihir0/hbapi"

// HotEntryStrategy represents using hot entry bookmark strategy.
type HotEntryStrategy struct {
	defaultStrategy
}

// NewHotEntryStrategy initialize HotEntryStrategy.
func NewHotEntryStrategy() HotEntryStrategy {
	return HotEntryStrategy{newDefaultStrategy("")}
}

// SetUser set user.
func (s *HotEntryStrategy) SetUser(user string) {
	panic("HotEntryStrategy: can not use user parameter.")
}

// SetTotalPages set totalPages.
func (s *HotEntryStrategy) SetTotalPages(pages int) {
	panic("HotEntryStrategy: can not use totalPages parameter.")
}

// SetMaxParallelRequest set maxParallelRequest.
func (s *HotEntryStrategy) SetMaxParallelRequest(max int) {
	panic("HotEntryStrategy: can not use maxParallelRequest parameter.")
}

func (s HotEntryStrategy) searchNeighbors() Neighbors {
	bookmarks := s.getHotEntryBookmarks()
	neighbors := s.calcNeighbors(bookmarks)
	return neighbors
}

func (s HotEntryStrategy) getHotEntryBookmarks() []string {
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
