package main

import (
	"fmt"
)

// reader is an interface that defines the act of reading data.
type reader interface {
	read(b []byte) (int, error)
}

// file defines a system file
type file struct {
	name string
}

// read implements the reader interface for a file.
func (file) read(b []byte) (int, error) {
	s := "<rss><channel><tittle>Going Go Programming</rss></channel></tittle>"
	copy(b, s)
	return len(s), nil
}

// pipe defines a named pipe network connection
type pipe struct {
	name string
}

// read implements the reader interface for a network connection.
func (pipe) read(b []byte) (int, error) {
	s := `{name: "bill", title: "developer"}`
	copy(b, s)
	return len(s), nil
}

// retrieve can read any device and process the data.
func retrieve(r reader) error {
	data := make([]byte, 100)

	len, err := r.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}

func main() {
	// Create two values one of type file nad one of type pipe.
	f := file{"data.json"}
	p := pipe{"cfg_service"}

	// Call the retrieve function for each conctrete type.
	retrieve(f)
	retrieve(p)
}
