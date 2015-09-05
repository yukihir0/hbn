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
	if n[i].Similarity == n[j].Similarity {
		ret = n[i].User < n[j].User
	} else {
		ret = n[i].Similarity > n[j].Similarity
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
