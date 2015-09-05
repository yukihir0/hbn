# hbn

"hbn" is library for search neighbors at hatena bookmark.

## Install

```
go get github.com/yukihir0/hbn
```

## How to use

### search neighbors by self bookmark.

```
client := hbn.NewClient()

user := "yukihir0"
bookmarks := client.RequestBookmarks(user)
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
```

### search neighbors by favorite bookmark.

```
client := hbn.NewClient()

user := "yukihir0"
bookmarks := client.RequestFavoriteBookmarks(user)
neighbors := client.SearchNeighbors(bookmarks)
...
```

### search neighbors by hot entry bookmark.

```
client := hbn.NewClient()

bookmarks := client.RequestHotEntryBookmarks()
neighbors := client.SearchNeighbors(bookmarks)
...
```

### search neighbors by search bookmark.

```
client := hbn.NewClient()

query := "golang"
bookmarks := client.RequestSearchBookmarks(query)
neighbors := client.SearchNeighbors(bookmarks)
...
```

### search neighbors by related bookmark.

```
client := hbn.NewClient()

user := "yukihir0"
bookmarks := client.RequestRelatedBookmarks(user)
neighbors := client.SearchNeighbors(bookmarks)
...
```

## License

Copyright &copy; 2015 yukihir0
