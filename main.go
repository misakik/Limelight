package main

import (
    "fmt"
    //"strconv"
    "github.com/blevesearch/bleve"
    "path/filepath"
    "os"
    "flag"
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
    if _, err := os.Stat(IndexDir); err == nil {
      os.RemoveAll(IndexDir)
    }
    mapping := bleve.NewIndexMapping()
    index, err := bleve.New(IndexDir, mapping)
    root := flag.Arg(1)
    ero := filepath.Walk(root,
      func(path string, f os.FileInfo, err error) error {
        count += 1
        size := f.Size()
        fmt.Printf("%d : Name: %s : Size: %d \n", count, path, size)
        data := Data{ Name: path, Size: size }
        index.Index(path, data)
        return nil
      })
    if ero != nil {
        fmt.Println(ero)
        return
    }

    fmt.Printf("filepath.Walk() returned %v\n", err)

  case "search":
    index, err := bleve.Open(IndexDir)
    if err != nil {
        fmt.Println(err)
        return
    }
    query := bleve.NewMatchQuery(flag.Arg(1))
    search := bleve.NewSearchRequest(query)
    searchResults, err := index.Search(search)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(searchResults)
  }

}
