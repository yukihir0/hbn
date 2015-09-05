package main

import (
	"fmt"

	"github.com/yukihir0/hbn"
)

func main() {
	client := hbn.NewClient()
	client.SetTotalPages(2)
	client.SetMaxParallelRequest(10)

	user := "yukihir0"
	bookmarks := client.RequestBookmarks(user)
	//bookmarks := client.RequestFavoriteBookmarks(user)
	//bookmarks := client.RequestHotEntryBookmarks()
	//bookmarks := client.RequestSearchBookmarks("golang")
	//bookmarks := client.RequestRelatedBookmarks(user)

	neighbors := client.SearchNeighbors(bookmarks)
	excluded := neighbors.Exclude([]string{
		user,
	})
	top := excluded.Top(20)

	for _, neighbor := range top {
		fmt.Printf(
			"[%s] : %.1f%% (%d/%d)\n",
			neighbor.User,
			neighbor.Similarity*100,
			len(neighbor.CommonBookmarks),
			len(bookmarks),
		)

		for _, bookmark := range neighbor.CommonBookmarks {
			fmt.Printf(" - %s\n", bookmark.Title)
		}
		fmt.Println()
	}
}
