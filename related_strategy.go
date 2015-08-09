package hbn

// RelatedStrategy represents using related bookmark strategy.
type RelatedStrategy struct {
	defaultStrategy
}

// NewRelatedStrategy initialize RelatedStrategy.
func NewRelatedStrategy(user string) RelatedStrategy {
	return RelatedStrategy{newDefaultStrategy(user)}
}

func (s RelatedStrategy) searchNeighbors() Neighbors {
	bookmarks := s.getRelatedBookmarks()
	neighbors := s.calcNeighbors(bookmarks)
	return neighbors
}

func (s RelatedStrategy) getRelatedBookmarks() []string {
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
