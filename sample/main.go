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
	strategy.SetTotalPages(10)
	//strategy.SetMaxParallelRequest(10)

	neighbors := hbn.SearchNeighbors(strategy)
	excluded := neighbors.Exclude([]string{user})
	top := excluded.Top(20)

	for _, neighbor := range top {
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
}
