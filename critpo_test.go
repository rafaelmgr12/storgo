package main

import (
	"bytes"
	"testing"
)

func TestCopyEncryptDecrypt(t *testing.T) {
	payload := "Foo not bar"
	src := bytes.NewReader([]byte(payload))
	dst := new(bytes.Buffer)
	key := newEncryptionKey()
	_, err := copyEncrypt(key, src, dst)
	if err != nil {
		t.Error(err)
	}

	out := new(bytes.Buffer)
	if _, err := copyDecrypt(key, dst, out); err != nil {
		t.Error(err)
	}

	if out.String() != payload {
		t.Errorf("decryption failed!!!")

	}

}
