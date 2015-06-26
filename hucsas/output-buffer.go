package main

import (
  "os"
  "bytes"
  )

type Buffer struct {
  buffer *bytes.Buffer
}

// Init initializes the Buffer instance.
// This should be called immediately after
// creating the object
func (b *Buffer) Init() {
  b.buffer = new(bytes.Buffer)
}

// WriteBufferToOutput writes the buffer to a files
// of the given name, replacing any existing file
func (b *Buffer) WriteBufferToOutput(filename string) error {
  fp, err := os.Create(filename)
  check(err)
  fp.Write(b.buffer.Bytes())
  fp.Close()
  return nil
}

// WriteByte appends the value to the buffer
func (b *Buffer) WriteByte(value byte) {
  b.buffer.WriteByte(value)
}

func (b *Buffer) Write(p []byte) (int, error) {
  return b.buffer.Write(p)
}
