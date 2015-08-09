// Package hbn : library for search neighbors at hatena bookmark.
package hbn

// SearchNeighbors search hatena bookmark neighbors.
func SearchNeighbors(s strategy) Neighbors {
	neighbors := s.searchNeighbors()
	return neighbors
}
