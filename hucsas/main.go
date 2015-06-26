package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
  )

const OUTPUT_VERSION uint8 = 1
const FILE_HEADER    string = "HXE"

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

    outputBuffer := new(Buffer)
    outputBuffer.Init()
    WriteOutputHeader(outputBuffer)

    for _, line := range lines {
      if strings.TrimLeft(line, " ")[0] != '#' {
        fmt.Println(line)
      }
    }

    check(outputBuffer.WriteBufferToOutput(opts.outFilename))
  }
}

func WriteOutputHeader(buf *Buffer) {
  for _, c := range FILE_HEADER {
    buf.WriteByte(byte(c))
  }
  buf.WriteByte(OUTPUT_VERSION)
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
