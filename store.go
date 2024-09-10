package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const defaultRootFolderName = "rrnetwork"

func CASPathTansformFunc(key string) PathKey {

	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blocksize := 5
	sliceLen := len(hashStr) / blocksize

	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blocksize, (i*blocksize)+blocksize
		paths[i] = hashStr[from:to]
	}

	return PathKey{
		PathName: strings.Join(paths, "/"),
		Filename: hashStr,
	}
}

type PathTransformFunc func(string) PathKey

type StoreOpts struct {
	// Root is the folder name of the root, containing all the folders of the system.
	Root              string
	PathTransformFunc PathTransformFunc
}

type PathKey struct {
	PathName string
	Filename string
}

func (p PathKey) FullPath() string {
	return fmt.Sprintf("%s/%s", p.PathName, p.Filename)
}

var DefaultPathTansformFunc = func(key string) PathKey {
	return PathKey{
		PathName: key,
		Filename: key,
	}
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	if opts.PathTransformFunc == nil {
		opts.PathTransformFunc = DefaultPathTansformFunc
	}

	if len(opts.Root) == 0 {
		opts.Root = defaultRootFolderName
	}

	return &Store{
		StoreOpts: opts,
	}
}

func (p PathKey) FirstPathName() string {
	paths := strings.Split(p.PathName, "/")
	if len(paths) == 0 {
		return ""
	}

	return paths[0]

}

func (s *Store) Has(key string) bool {

	pathKey := s.PathTransformFunc(key)

	fullPathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())
	// remove : from fullPathWithRoot
	fullPathWithRoot = strings.Replace(fullPathWithRoot, ":", "", -1)

	_, err := os.Stat(fullPathWithRoot)
	return !os.IsNotExist(err)

}

func (s *Store) Clear() error {
	return os.RemoveAll(s.Root)
}

func (s *Store) Write(key string, r io.Reader) (int64, error) {
	return s.writeStream(key, r)
}

func (s *Store) Delete(key string) error {
	pathKey := s.PathTransformFunc(key)

	defer func() {
		log.Printf("deleted [%s] from disk", pathKey.Filename)
	}()

	firstPathNameWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FirstPathName())

	firstPathNameWithRoot = strings.Replace(firstPathNameWithRoot, ":", "", -1)

	return os.RemoveAll(firstPathNameWithRoot)
}

func (s *Store) Read(key string) (int64, io.Reader, error) {

	return s.readStream(key)
}

func (s *Store) readStream(key string) (int64, io.ReadCloser, error) {

	pathKey := s.PathTransformFunc(key)
	fullPathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())
	// remove : from fullPathWithRoot
	fullPathWithRoot = strings.Replace(fullPathWithRoot, ":", "", -1)

	file, err := os.Open(fullPathWithRoot)
	if err != nil {
		return 0, nil, err
	}

	fi, err := os.Stat(fullPathWithRoot)
	if err != nil {
		return 0, nil, err
	}

	return fi.Size(), file, nil
}

func (s *Store) writeStream(key string, r io.Reader) (int64, error) {
	pathKey := s.PathTransformFunc(key)

	pathNameWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.PathName)
	// remove : from pathNameWithRoot
	pathNameWithRoot = strings.Replace(pathNameWithRoot, ":", "", -1)
	if err := os.MkdirAll(pathNameWithRoot, os.ModePerm); err != nil {
		return 0, err
	}

	fullPathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())
	// remove : from fullPathWithRoot
	fullPathWithRoot = strings.Replace(fullPathWithRoot, ":", "", -1)

	f, err := os.Create(fullPathWithRoot)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n, err := io.Copy(f, r)

	if err != nil {
		return 0, err
	}

	log.Printf("written (%d) bytes to disk: %s", n, fullPathWithRoot)

	return n, nil
}
