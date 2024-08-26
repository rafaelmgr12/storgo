package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathname := CASPathTansformFunc(key)
	expectedOriginalKey := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	if pathname.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathname.PathName, expectedPathName)

	}

	if pathname.Filename != expectedOriginalKey {
		t.Errorf("have %s want %s", pathname.Filename, expectedOriginalKey)

	}

}

func TestStore(t *testing.T) {
	s := newStore()
	defer teardown(t, s)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("foo_%d", i)

		data := []byte("some jpg bytes")

		_, err := s.writeStream(key, bytes.NewReader(data))
		if err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); !ok {
			t.Errorf("expetected to have key %s", key)
		}

		r, err := s.Read(key)

		if err != nil {
			t.Error(err)
		}

		b, _ := io.ReadAll(r)

		if string(b) != string(data) {
			t.Errorf("want %s have %s", data, b)
		}

		if err := s.Delete(key); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); ok {
			t.Errorf("expected to NOT have key %s", key)
		}
	}

}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTansformFunc,
	}

	return NewStore(opts)
}

func teardown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
