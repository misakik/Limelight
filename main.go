package main

import (
    "bytes"
    "fmt"
    "path/filepath"
    "os"
    "flag"
    "net/http"
    "io"
    "io/ioutil"
    "mime/multipart"
    "github.com/blevesearch/bleve"
)

var count int = 0
type Data struct {
  Name string
  Size int64
  Text string
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
    err = filepath.Walk(root,
      func(path string, f os.FileInfo, err error) error {
        count += 1

        size := f.Size()
        text := ""

        if !f.IsDir() && size < 10000000 {
          client := &http.Client{}

          bodyBuf := &bytes.Buffer{}
          bodyWriter := multipart.NewWriter(bodyBuf)

          fileWriter, err := bodyWriter.CreateFormFile("uploadfile", path)
          if err != nil {
              fmt.Println("error writing to buffer")
              return err
          }

          fh, err := os.Open(path)
          if err != nil {
              fmt.Println("error opening file")
              return err
          }
          defer fh.Close()

          _, err = io.Copy(fileWriter, fh)
          if err != nil {
              return err
          }

          //contentType := bodyWriter.FormDataContentType()
          bodyWriter.Close()

          url := "http://localhost:9998/tika"
          request, err := http.NewRequest("PUT", url, bodyBuf)
          if err != nil {
            fmt.Println(err)
            return nil
          }

          request.Header.Set("Accept", "text/plain")

          response, err := client.Do(request)
          if err != nil {
            fmt.Println(err)
            return nil
          }

          defer response.Body.Close()
          body, err := ioutil.ReadAll(response.Body)
          if err != nil {
            fmt.Println(err)
            return nil
          }
          text = string(body)

        }

        fmt.Printf("%d : Name: %s : Size: %d : Text: %s \n", count, path, size, text)
        data := Data{ Name: path, Size: size, Text: text}
        index.Index(path, data)
        return nil
      })
    if err != nil {
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
