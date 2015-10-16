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

func main() {
  flag.Parse()

  switch flag.Arg(0) {
  case "index" :
    mapping := bleve.NewIndexMapping()
    index, err := bleve.New("example.bleve", mapping)
    root := flag.Arg(1)
    ero := filepath.Walk(root,
      func(path string, f os.FileInfo, err error) error {
        fmt.Printf("Visited: %s : %d \n", path, count)
        count += 1
        data := struct {
          Name string
        }{
          Name: path,
        }
        index.Index(path, data)
        return nil
      })
    if ero != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("filepath.Walk() returned %v\n", err)
    fmt.Printf("%d", count)
    
  case "search":
    index, err := bleve.Open("example.bleve")
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
    fmt.Println(index)
  }

}
