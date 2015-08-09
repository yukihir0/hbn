package hbn

// HotEntryStrategy represents using hot entry bookmark strategy.
type HotEntryStrategy struct {
	defaultStrategy
}

// NewHotEntryStrategy initialize HotEntrystrategy.
func NewHotEntryStrategy() HotEntryStrategy {
	return HotEntryStrategy{newDefaultStrategy("")}
}

func (s HotEntryStrategy) searchNeighbors() Neighbors {
	bookmarks := s.getHotEntryBookmarks()
	neighbors := s.calcNeighbors(bookmarks)
	return neighbors
}
