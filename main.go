package main

import (
  "flag"
  "fmt"
)

func main() {
  addLicense := NewAddLicense()

  license := flag.String("l", "GPLv3", "License to add.")
  outputDir := flag.String("o", "", "Where to write the license. If not provided, current working directory will be used.")
  list := flag.Bool("list", false, "List Supported Licenses")

  flag.Parse()

  err := addLicense.Init()
  if err != nil {
    fmt.Println(err)
    return
  }

  if *list == true {
    licenses := addLicense.ListLicenses()
    fmt.Println(licenses)
    return;
  }

  err = addLicense.Add(*license, *outputDir)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println("Done")
}
