package main

import (
  "os"
  "bytes"
  )

type Buffer struct {
  buffer *bytes.Buffer
}

func (b *Buffer) Init() {
  b.buffer = new(bytes.Buffer)
}

func (b *Buffer) WriteBufferToOutput(filename string) error {
  fp, err := os.Create(filename)
  check(err)
  fp.Write(b.buffer.Bytes())
  fp.Close()
  return nil
}

func (b *Buffer) WriteByte(value byte) {
  b.buffer.WriteByte(value)
}
