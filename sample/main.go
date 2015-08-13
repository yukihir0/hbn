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
	//strategy := hbn.NewFavoriteStrategy(user)
	//strategy := hbn.NewSearchStrategy("golang")
	//strategy.SetUser(user)
	//strategy.SetTotalPages(2)
	//strategy.SetMaxParallelRequest(10)

	neighbors := hbn.SearchNeighbors(strategy)

	limit := 20
	for _, neighbor := range neighbors {
		if neighbor.User != user {
			fmt.Printf(
				"[%s] : %.1f%% (%d/%d)\n",
				neighbor.User,
				neighbor.GetSimilarity()*100,
				neighbor.GetCommonBookmarkCount(),
				neighbor.GetAllBookmarkCount(),
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
