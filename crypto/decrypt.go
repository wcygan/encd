package crypto

import (
	"os"
)

func Decrypt(ciphertext []byte, oracle *Oracle, out *os.File) error {
	defer out.Close()

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
