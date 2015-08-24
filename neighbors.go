package hbn

// Neighbors represents list of Neighbor.
type Neighbors []Neighbor

func (n Neighbors) Len() int {
	return len(n)
}

func (n Neighbors) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Neighbors) Less(i, j int) bool {
	var ret bool
	if len(n[i].CommonBookmarks) == len(n[j].CommonBookmarks) {
		ret = n[i].User < n[j].User
	} else {
		ret = len(n[i].CommonBookmarks) > len(n[j].CommonBookmarks)
	}
	return ret
}

func (n Neighbors) Exclude(users []string) Neighbors {
	excluded := Neighbors{}
	for _, neighbor := range n {
		if !neighbor.included(users) {
			excluded = append(excluded, neighbor)
		}
	}
	return excluded
}

func (n Neighbors) Top(max int) Neighbors {
	return n[0:max]
}
