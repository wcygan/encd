package crypto

import (
	"io"
)

func Encrypt(plaintext []byte, oracle *Oracle, out io.Writer) error {
	ciphertext := oracle.Encrypt(plaintext)

	_, err := out.Write(ciphertext)
	if err != nil {
		return err
	}

	return nil
}
