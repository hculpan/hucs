package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
  )

const HUCS_ASSEMBLER_VERSION string = "0.1"
const OUTPUT_VERSION          uint8 = 1
const FILE_HEADER            string = "HXE"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// main
func main() {
  fmt.Printf("Hucs Assembler v%s\n", HUCS_ASSEMBLER_VERSION)

  if opts, err := ParseCommandLine(); err != nil {
    fmt.Println(err)
    Usage()
  } else {
    if opts.verbose {
      fmt.Println("   input=" + opts.inFilename)
      fmt.Println("  output=" + opts.outFilename)
    }

    lines, err := ReadLines(opts.inFilename)
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

// WriteOutputHeader writes the file magic number ("HXE")
// and current vm file version to the buffer
func WriteOutputHeader(buf *Buffer) {
  for _, c := range FILE_HEADER {
    buf.WriteByte(byte(c))
  }
  buf.WriteByte(OUTPUT_VERSION)
}

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
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
