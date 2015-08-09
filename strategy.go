package hbn

type strategy interface {
	searchNeighbors() Neighbors
}
