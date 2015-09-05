package hbn

// Neighbor represents hatena bookmark neighbor.
type Neighbor struct {
	User            string
	Similarity      float64
	CommonBookmarks []Bookmark
}

// NewDummyNeighbor return dummy of neighbor.
func NewDummyNeighbor(user string, similarity float64) Neighbor {
	dummy := Neighbor{
		User:            user,
		Similarity:      similarity,
		CommonBookmarks: []Bookmark{},
	}
	return dummy
}

func (n Neighbor) included(users []string) bool {
	for _, user := range users {
		if user == n.User {
			return true
		}
	}
	return false
}
