package hbn

import "github.com/yukihir0/hbapi"

// SearchStrategy represents using search bookmark strategy.
type SearchStrategy struct {
	defaultStrategy
	query string
}

// NewSearchStrategy initialize SearchStrategy.
func NewSearchStrategy(query string) SearchStrategy {
	s := SearchStrategy{
		defaultStrategy: newDefaultStrategy(""),
	}
	s.SetQuery(query)
	return s
}

// SetUser set user.
func (s *SearchStrategy) SetUser(user string) {
	panic("SearchStrategy: can not use user parameter.")
}

// SetTotalPages set totalPages.
func (s *SearchStrategy) SetTotalPages(pages int) {
	panic("SearchStrategy: can not use totalPages parameter.")
}

// SetMaxParallelRequest set maxParallelRequest.
func (s *SearchStrategy) SetMaxParallelRequest(max int) {
	panic("SearchStrategy: can not use maxParallelRequest parameter.")
}

func (s *SearchStrategy) SetQuery(query string) {
	if query != "" {
		s.query = query
	}
}

func (s SearchStrategy) searchNeighbors() Neighbors {
	bookmarks := s.getSearchBookmarks()
	neighbors := s.calcNeighbors(bookmarks)
	return neighbors
}

func (s SearchStrategy) getSearchBookmarks() []string {
	bookmarks := []string{}

	params := hbapi.NewSearchFeedParams(s.query)
	feed, err := hbapi.GetSearchFeed(params)
	if err != nil {
		return bookmarks
	}

	for _, item := range feed.Items {
		bookmarks = append(bookmarks, item.Link)
	}

	return bookmarks
}
