package hbn

type api interface {
	RequestBookmarks(user string, page int) []Bookmark
	RequestFavoriteBookmarks(user string, page int) []Bookmark
	RequestHotEntryBookmarks() []Bookmark
	RequestSearchBookmarks(query string) []Bookmark
	RequestRelatedBookmarks(url string) []Bookmark
	RequestUserBookmark(url string) map[string]Bookmark
}
