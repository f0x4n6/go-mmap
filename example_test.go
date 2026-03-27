package mmap_test

import (
	"fmt"
	"log"
	"os"

	"github.com/f0x4n6/go-mmap"
)

func ExampleMapRegion() {
	m, err := mmap.MapRegion(nil, 100, mmap.RDONLY, mmap.ANON, 0)
	if err != nil {
		log.Fatal(err)
	}
	// m acts as a writable slice of bytes that is not managed by the Go runtime.
	fmt.Println(len(m))

	// Because the region is not managed by the Go runtime, the Unmap method should
	// be called when finished with it to avoid leaking memory.
	if err := m.Unmap(); err != nil {
		log.Fatal(err)
	}

	// Output: 100
}

func ExampleMap() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString("Hello, world")
	if err != nil {
		log.Fatal(err)
	}
	// The file must be closed, even after calling Unmap.
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	m, err := mmap.Map(f, mmap.RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}

	// m acts as a writable slice of bytes that is a view into the open file, notes.txt.
	// It is sized to the file contents automatically.
	fmt.Println(string(m))

	// The Unmap method should be called when finished with it to avoid leaking memory
	if err := m.Unmap(); err != nil {
		log.Fatal(err)
	}

	// Hello, world
}
