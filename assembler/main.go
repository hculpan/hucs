package main

import (
  "fmt"
  "errors"
  "flag"
  )

func main() {
  fmt.Println("Hucs Assembler")

  outFilenamePtr := flag.String("of", nil, "Output filename")
  flag.Parse()
  fmt.Println("outFilename=" + *outFilenamePtr)
  fmt.Println(flag.Args())
}

func usage() {
  fmt.Println("usage: hucsas [-of=outputfile] <input filename>")
  fmt.Println("-of    set output filename")
}
