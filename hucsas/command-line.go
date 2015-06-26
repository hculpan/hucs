package main

import (
  "fmt"
  "flag"
  "errors"
  "strings"
  )

type Options struct {
  outFilename string
  inFilename string
}

func buildOutFilename(inFilename string) string {
  if strings.HasSuffix(inFilename, ".as") {
    return inFilename[:len(inFilename)-3] + ".hxe"
  } else {
    return inFilename + ".hxe"
  }
}

func parseCommandLine() (*Options, error) {
  options := Options{}

  outFilenamePtr := flag.String("of", "", "Output filename")

  var outFilename string
  var inFilename string

  flag.Parse()
  if len(flag.Args()) < 1 {
    return nil, errors.New("Missing required parameters")
  } else if len(flag.Args()) > 1 {
    return nil, errors.New("Too many parameters given")
  } else {
    inFilename = flag.Args()[0]
    if *outFilenamePtr == "" {
      outFilename = buildOutFilename(inFilename)
    } else {
      outFilename = *outFilenamePtr
    }

    options.outFilename = outFilename
    options.inFilename = inFilename
  }

  return &options, nil
}

func usage() {
  fmt.Println("usage: hucsas [-of=outputfile] <input filename>")
  fmt.Println("    -of    set output filename")
}
