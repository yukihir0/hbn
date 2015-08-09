package hbn

// SelfStrategy represents using self bookmark strategy.
type SelfStrategy struct {
	defaultStrategy
}

// NewSelfStrategy initialize SelfStrategy.
func NewSelfStrategy(user string) SelfStrategy {
	return SelfStrategy{newDefaultStrategy(user)}
}

func (s SelfStrategy) searchNeighbors() Neighbors {
	bookmarks := s.getBookmarks()
	neighbors := s.calcNeighbors(bookmarks)
	return neighbors
}
