package main

import (
    "fmt"
    "../utils/io"
)

func main() {
   filepath := "/Users/amit.upadhyay/Documents/GoPrograms/aku/utils/io/fileutils.go"
   lis, err := io.ReadFile(filepath, false)
   if err != nil {
        fmt.Println(err)
   }
   for _, l := range lis {
        fmt.Println(l)
   }
}
