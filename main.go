package main

import (
    "fmt"
    "path/filepath"
    "os"
    "flag"
    "github.com/blevesearch/bleve"
)

var count int = 0
type Data struct {
  Name string
  Size int64
}
const IndexDir = ".tmp/index.data"

func main() {
  flag.Parse()

  switch flag.Arg(0) {
  case "index" :
    // Delete index directory if it already exists
    if _, err := os.Stat(IndexDir); err == nil {
      os.RemoveAll(IndexDir)
    }

    mapping := bleve.NewIndexMapping()
    index, err := bleve.New(IndexDir, mapping)
    if err != nil {
        fmt.Println(err)
        return
    }

    root := flag.Arg(1)
    error := filepath.Walk(root,
      func(path string, f os.FileInfo, err error) error {
        count += 1
        size := f.Size()
        fmt.Printf("%d : Name: %s : Size: %d \n", count, path, size)
        data := Data{ Name: path, Size: size }
        index.Index(path, data)
        return nil
      })
    if error != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Index Done. %d items.\n", count)

  case "search":
    index, err := bleve.Open(IndexDir)
    if err != nil {
        fmt.Println(err)
        return
    }
    query := bleve.NewMatchQuery(flag.Arg(1))
    searchRequest := bleve.NewSearchRequest(query)
    searchResults, err := index.Search(searchRequest)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(searchResults)
  }

}
