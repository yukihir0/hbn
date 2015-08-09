# hbn

"hbn" is library for search neighbors at hatena bookmark.

## Install

```
go get github.com/yukihir0/hbn
```

## How to use

### search neighbors by self bookmark.

```
user := "yukihir0"
strategy := hbn.NewSelfStrategy(user)
neighbors := hbn.SearchNeighbors(strategy)

for _, neighbor := range neighbors {
  fmt.Printf(
    "[%s] : %.1f%% (%d/%d)\n",
    neighbor.User,
    neighbor.GetSimilarity()*100,
    neighbor.GetCommonBookmarkCount(),
    neighbor.GetTotalBookmarkCount(),
  )
}
```

### search neighbors by related bookmark.

```
user := "yukihir0"
strategy := hbn.NewRelatedStrategy(user)
neighbors := hbn.SearchNeighbors(strategy)

for _, neighbor := range neighbors {
  fmt.Printf(
    "[%s] : %.1f%% (%d/%d)\n",
    neighbor.User,
    neighbor.GetSimilarity()*100,
    neighbor.GetCommonBookmarkCount(),
    neighbor.GetTotalBookmarkCount(),
  )
}
```

### search neighbors by hot entry bookmark.

```
strategy := hbn.NewHotEntryStrategy()
neighbors := hbn.SearchNeighbors(strategy)

for _, neighbor := range neighbors {
  fmt.Printf(
    "[%s] : %.1f%% (%d/%d)\n",
    neighbor.User,
    neighbor.GetSimilarity()*100,
    neighbor.GetCommonBookmarkCount(),
    neighbor.GetTotalBookmarkCount(),
  )
}
```

## License

Copyright &copy; 2015 yukihir0
