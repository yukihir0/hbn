package hbn

import "github.com/yukihir0/hbapi"

// Neighbor represents hatena bookmark neighbor.
type Neighbor struct {
	User            string
	CommonBookmarks []hbapi.Entry
	AllBookmarks    []hbapi.Entry
}

// GetCommonBookmarkCount return common bookmark count.
func (n Neighbor) GetCommonBookmarkCount() int {
	return len(n.CommonBookmarks)
}

// GetAllBookmarkCount return all bookmark count.
func (n Neighbor) GetAllBookmarkCount() int {
	return len(n.AllBookmarks)
}

// GetSimilarity return similarity of neighbor.
func (n Neighbor) GetSimilarity() float64 {
	return float64(n.GetCommonBookmarkCount()) / float64(n.GetAllBookmarkCount())
}

func (n Neighbor) included(users []string) bool {
	for _, user := range users {
		if user == n.User {
			return true
		}
	}
	return false
}
