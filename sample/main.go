package main

import (
	"fmt"

	"github.com/yukihir0/hbn"
)

func main() {
	user := "yukihir0"
	strategy := hbn.NewSelfStrategy(user)
	//strategy := hbn.NewRelatedStrategy(user)
	//strategy := hbn.NewHotEntryStrategy()
	//strategy.User = user
	//strategy.TotalPages = 2
	//strategy.MaxParallelRequest = 20

	neighbors := hbn.SearchNeighbors(strategy)

	limit := 20
	for _, neighbor := range neighbors {
		if neighbor.User != user {
			fmt.Printf(
				"[%s] : %.1f%% (%d/%d)\n",
				neighbor.User,
				neighbor.GetSimilarity()*100,
				neighbor.GetCommonBookmarkCount(),
				neighbor.GetTotalBookmarkCount(),
			)
			for _, entry := range neighbor.CommonBookmarks {
				fmt.Printf(" - %s\n", entry.Title)
			}
			fmt.Println()
		}
		limit--
		if limit <= 0 {
			break
		}
	}
}
