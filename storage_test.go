package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathname := CASPathTransformFunc(key)
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	expectedOriginal := "6804429f74181a63c50c3d81d733a12f14a353ff"
	fmt.Println(pathname)

	if pathname.PathName != expectedPathName {
		t.Errorf("have %s, want %s", pathname.PathName, expectedPathName)
	}

	if pathname.Filename != expectedOriginal {
		t.Errorf("have %s, want %s", pathname.Filename, expectedOriginal)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)
	key := "myspecialpicture"
	data := []byte("some data")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Errorf("Error writing stream: %s", err)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Errorf("Error reading stream: %s", err)
	}
	b, err := io.ReadAll(r)
	if err != nil {
		t.Errorf("Error reading all: %s", err)
	}
	if string(b) != string(data) {
		t.Errorf("have %s, want %s", b, data)
	}

	fmt.Println(string(b))
}
