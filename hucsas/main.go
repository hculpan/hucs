package main

import (
  "fmt"
  "bufio"
  "os"
  "bytes"
  "unsafe"
  )

const HUCS_ASSEMBLER_VERSION  string = "0.1"
const VM_VERSION              uint8 = 1
const MAGIC_NUMBER            string = "HXE"

type HxeHeader struct {
  version byte
  codeOffset uint32
  codeSize uint32
  dataOffset uint32
  dataSize uint32
  exportOffset uint32
  exportSize uint32
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// main
func main() {
  var wasError bool = false
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

    if opts.verbose {
      fmt.Printf("Input source: lines=%d\n", len(lines))
    }

    for i, line := range lines {
      lineBuf, err := AssembleLine(line)
      if err != nil {
        fmt.Printf("Error: %s [%d]\n", err, i)
        wasError = true
      } else if opts.verbose {
        OutputAssembledLine(line, lineBuf)
      }

      outputBuffer.Write(lineBuf.Bytes())
    }

    if !wasError {
      check(outputBuffer.WriteBufferToOutput(opts.outFilename))
    } else {
      fmt.Println("Errors found, process failed")
    }
  }
}

func OutputAssembledLine(line string, lineBuf *bytes.Buffer) {
  if len(lineBuf.Bytes()) == 0 {
    fmt.Printf("    ")
  } else {
    for _, b := range lineBuf.Bytes() {
      fmt.Printf(" $%02X", b)
    }
  }
  fmt.Printf(" : %s\n", line)
}

// WriteOutputHeader writes the file magic number ("HXE")
// and current vm file version to the buffer
func WriteOutputHeader(buf *Buffer) {
  for _, c := range MAGIC_NUMBER {
    buf.WriteByte(byte(c))
  }

  hxeHeader := new(HxeHeader)
  hxeHeader.version = VM_VERSION
  hxeHeader.dataOffset = 3 + uint32(unsafe.Sizeof(hxeHeader))
  hxeHeader.dataSize = 0
//  hxeHeader.exportOffset = 
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
