package main

import (
  "fmt"
  "flag"
  "strings"
  "errors"
  "bufio"
  "os"
  "bytes"
  )

const OUTPUT_VERSION uint8 = 1
const FILE_HEADER    string = "HXE"

type Options struct {
  outFilename string
  inFilename string
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
  fmt.Println("Hucs Assembler")

  if opts, err := parseCommandLine(); err != nil {
    fmt.Println(err)
    usage()
  } else {
    fmt.Println("input=" + opts.inFilename)
    fmt.Println("output=" + opts.outFilename)

    lines, err := readLines(opts.inFilename)
    check(err);

    outputBuffer := writeOutputHeader()

    for _, line := range lines {
      if strings.TrimLeft(line, " ")[0] != '#' {
        fmt.Println(line)
      }
    }

    check(writeBufferToOutput(outputBuffer, opts.outFilename))
  }
}

func writeBufferToOutput(buf *bytes.Buffer, filename string) error {
  fp, err := os.Create(filename)
  check(err)
  fp.Write(buf.Bytes())
  fp.Close()
  return nil
}

func writeOutputHeader() *bytes.Buffer {
  buf := new(bytes.Buffer)
  for _, c := range FILE_HEADER {
    buf.WriteByte(byte(c))
  }
  buf.WriteByte(OUTPUT_VERSION)

  return buf
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
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
