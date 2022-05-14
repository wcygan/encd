package crypto

import (
	"io"
)

func Decrypt(ciphertext []byte, oracle *Oracle, out io.Writer) error {
	plaintext, err := oracle.Decrypt(ciphertext)
	if err != nil {
		return err
	}

	_, err = out.Write(plaintext)
	if err != nil {
		return err
	}

	return nil
}
