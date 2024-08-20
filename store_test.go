package main

import (
	"bytes"
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

	if pathname.Original != expectedOriginalKey {
		t.Errorf("have %s want %s", pathname.Original, expectedOriginalKey)

	}

}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTansformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream("myspecialpictures", data); err != nil {
		t.Error(err)
	}

}
