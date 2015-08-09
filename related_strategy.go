package hbn

// RelatedStrategy represents using related bookmark strategy.
type RelatedStrategy struct {
	defaultStrategy
}

// NewRelatedStrategy initialize Relatedstrategy.
func NewRelatedStrategy(user string) RelatedStrategy {
	return RelatedStrategy{newDefaultStrategy(user)}
}

func (s RelatedStrategy) searchNeighbors() Neighbors {
	bookmarks := s.getRelatedBookmarks()
	neighbors := s.calcNeighbors(bookmarks)
	return neighbors
}
