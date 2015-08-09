package hbn

import "github.com/yukihir0/hbapi"

// Neighbor represents hatena bookmark neighbor.
type Neighbor struct {
	User            string
	CommonBookmarks []hbapi.Entry
	TotalBookmarks  []hbapi.Entry
}

// GetCommonBookmarkCount return common bookmark count.
func (n Neighbor) GetCommonBookmarkCount() int {
	return len(n.CommonBookmarks)
}

// GetTotalBookmarkCount return total bookmark count.
func (n Neighbor) GetTotalBookmarkCount() int {
	return len(n.TotalBookmarks)
}

// GetSimilarity return similarity of neighbor.
func (n Neighbor) GetSimilarity() float64 {
	return float64(n.GetCommonBookmarkCount()) / float64(n.GetTotalBookmarkCount())
}
