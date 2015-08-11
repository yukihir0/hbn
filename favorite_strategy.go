package hbn

import "github.com/yukihir0/hbapi"

// FavoriteStrategy represents using favorite bookmark strategy.
type FavoriteStrategy struct {
	defaultStrategy
}

// NewFavoriteStrategy initialize FavoriteStrategy.
func NewFavoriteStrategy(user string) FavoriteStrategy {
	return FavoriteStrategy{newDefaultStrategy(user)}
}

func (s FavoriteStrategy) searchNeighbors() Neighbors {
	bookmarks := s.getFavoriteBookmarks()
	neighbors := s.calcNeighbors(bookmarks)
	return neighbors
}

func (s FavoriteStrategy) getFavoriteBookmarks() []string {
	bookmarksChan := make(chan []string)
	limitChan := make(chan struct{}, s.maxParallelRequest)

	go func() {
		for page := 0; page < s.totalPages; page++ {
			select {
			case limitChan <- empty:
				go func(user string, page int) {
					params := hbapi.NewFavoriteFeedParams(user)
					params.SetPage(page)
					feed, err := hbapi.GetFavoriteFeed(params)
					if err != nil {
						bookmarksChan <- []string{}
						<-limitChan
						return
					}

					urls := []string{}
					for _, item := range feed.Items {
						urls = append(urls, item.Link)
					}
					bookmarksChan <- urls
					<-limitChan
				}(s.user, page)
			}
		}
	}()

	bookmarks := []string{}
	for i := 0; i < s.totalPages; i++ {
		bookmarks = append(bookmarks, <-bookmarksChan...)
	}

	return bookmarks
}
