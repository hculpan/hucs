package main

import (
  "fmt"
  "flag"
  "errors"
  "strings"
  )

type Options struct {
  outFilename string
  inFilename  string
  verbose     bool
}

// BuildOutFilename takes the input filename and
// returns the same filename with with an ".hxe"
// extension.  If the file has the ".as" extension,
// it replaces that; otherwise it just adds the
// hxe extension
func BuildOutFilename(inFilename string) string {
  if strings.HasSuffix(inFilename, ".as") {
    return inFilename[:len(inFilename)-3] + ".hxe"
  } else {
    return inFilename + ".hxe"
  }
}

// ParseCommandLine parses the command line and
// returns an instance of the Options type.  If
// output filename is not specified on the command
// line then it populates this with the default, so
// both input and output filenames should be populates
// no matter what the user specifies...unless there's
// an error, of course.
func ParseCommandLine() (*Options, error) {
  options := Options{verbose: false}

  outFilenamePtr := flag.String("of", "", "Output filename")
  verbosePtr := flag.Bool("v", false, "verbose output")

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
      outFilename = BuildOutFilename(inFilename)
    } else {
      outFilename = *outFilenamePtr
    }

    options.outFilename = outFilename
    options.inFilename = inFilename
    options.verbose = *verbosePtr
  }

  return &options, nil
}

// Usage prints the command usage for the assembler
func Usage() {
  fmt.Println("usage: hucsas [-of=outputfile] <input filename>")
  fmt.Println("    -of    set output filename")
  fmt.Println("    -v     verbose output on")
}
