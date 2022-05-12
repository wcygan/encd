package encoder

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Encode(contents []byte, password string, out io.Writer) error {
	key := []byte(password)

	c, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	_, err = out.Write(gcm.Seal(nonce, nonce, contents, nil))
	if err != nil {
		return err
	}

	return nil
}
