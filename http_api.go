package hbn

import "github.com/yukihir0/hbapi"

type HttpAPI struct {
}

func NewHttpAPI() api {
	return HttpAPI{}
}

func (api HttpAPI) RequestBookmarks(user string, page int) []Bookmark {
	params := hbapi.NewFeedParams(user)
	params.SetPage(page)
	feed, err := hbapi.GetFeed(params)
	if err != nil {
		return []Bookmark{}
	}

	bookmarks := []Bookmark{}
	for _, item := range feed.Items {
		bookmarks = append(bookmarks, Bookmark{
			Title: item.Title,
			URL:   item.Link,
			Count: item.BookmarkCount,
		})
	}

	return bookmarks
}

func (api HttpAPI) RequestFavoriteBookmarks(user string, page int) []Bookmark {
	params := hbapi.NewFavoriteFeedParams(user)
	params.SetPage(page)
	feed, err := hbapi.GetFavoriteFeed(params)
	if err != nil {
		return []Bookmark{}
	}

	bookmarks := []Bookmark{}
	for _, item := range feed.Items {
		bookmarks = append(bookmarks, Bookmark{
			Title: item.Title,
			URL:   item.Link,
			Count: item.BookmarkCount,
		})
	}

	return bookmarks
}

func (api HttpAPI) RequestHotEntryBookmarks() []Bookmark {
	params := hbapi.NewHotEntryFeedParams()
	feed, err := hbapi.GetHotEntryFeed(params)
	if err != nil {
		return []Bookmark{}
	}

	bookmarks := []Bookmark{}
	for _, item := range feed.Items {
		bookmarks = append(bookmarks, Bookmark{
			Title: item.Title,
			URL:   item.Link,
			Count: item.BookmarkCount,
		})
	}

	return bookmarks
}

func (api HttpAPI) RequestSearchBookmarks(query string) []Bookmark {
	params := hbapi.NewSearchFeedParams(query)
	feed, err := hbapi.GetSearchFeed(params)
	if err != nil {
		return []Bookmark{}
	}

	bookmarks := []Bookmark{}
	for _, item := range feed.Items {
		bookmarks = append(bookmarks, Bookmark{
			Title: item.Title,
			URL:   item.Link,
			Count: item.BookmarkCount,
		})
	}

	return bookmarks
}

func (api HttpAPI) RequestRelatedBookmarks(url string) []Bookmark {
	info, err := hbapi.GetEntryInfo(url)
	if err != nil {
		return []Bookmark{}
	}

	bookmarks := []Bookmark{}
	for _, bookmark := range info.Related {
		bookmarks = append(bookmarks, Bookmark{
			Title: bookmark.Title,
			URL:   bookmark.URL,
			Count: bookmark.Count,
		})
	}

	return bookmarks
}

func (api HttpAPI) RequestUserBookmark(url string) map[string]Bookmark {
	info, err := hbapi.GetEntryInfo(url)
	if err != nil {
		return map[string]Bookmark{}
	}

	ub := map[string]Bookmark{}
	for _, bookmark := range info.Bookmarks {
		ub[bookmark.User] = Bookmark{
			Title: info.Entry.Title,
			URL:   info.Entry.URL,
			Count: info.Entry.Count,
		}
	}

	return ub
}
