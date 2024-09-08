package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func newEncryptionKey() []byte {
	keyBuff := make([]byte, 32)
	io.ReadFull(rand.Reader, keyBuff)
	return keyBuff
}

func copyDecrypt(key []byte, src io.Reader, dst io.Writer) (int, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return 0, err
	}

	// Read the IV from the given io.Reader, which in our case shoudl be the
	// the block.BlockSize() bytes we read.

	iv := make([]byte, block.BlockSize())
	if _, err := src.Read(iv); err != nil {
		return 0, nil
	}

	buf := make([]byte, 32*1024)
	stream := cipher.NewCTR(block, iv)

	for {
		n, err := src.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf, buf[:n])
			if _, err := dst.Write(buf[:n]); err != nil {
				return 0, err
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

	}

	return 0, nil

}

func copyEncrypt(key []byte, src io.Reader, dst io.Writer) (int, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return 0, err
	}

	iv := make([]byte, block.BlockSize()) // 16 bytes
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return 0, err
	}

	// prepend the IV to the file
	if _, err := dst.Write(iv); err != nil {
		return 0, err
	}

	buf := make([]byte, 32*1024)
	stream := cipher.NewCTR(block, iv)

	for {
		n, err := src.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf, buf[:n])
			if _, err := dst.Write(buf[:n]); err != nil {
				return 0, err
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, err
		}

	}
	return 0, nil

}
