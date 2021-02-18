package readers

import (
	"errors"
	"io/ioutil"
)

// CharReader is a way to read characters from an input
type CharReader struct {
	chars []byte
	index int
	size  int
}

// NewCharReader initialize a new character reader
func NewCharReader(path string) (reader CharReader) {
	chars, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	reader.chars = chars
	reader.index = 0
	reader.size = len(chars)
	return
}

// Consume a character from the file
func (reader *CharReader) Consume() byte {
	if reader.index < reader.size-1 {
		b := reader.chars[reader.index]
		reader.index++
		return b
	}

	panic(errors.New("end of input reached"))
}

// Peek the next character from the file
func (reader *CharReader) Peek() byte {
	if reader.index < reader.size-1 {
		return reader.chars[reader.index]
	}

	panic(errors.New("end of input reached"))
}

// IsEOF checks whether the the reader has reached the end of the file
func (reader *CharReader) IsEOF() bool {
	if reader.index == reader.size-1 {
		return true
	}

	return false
}
