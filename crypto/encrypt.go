package crypto

import (
	"os"
)

func Encrypt(plaintext []byte, oracle *Oracle, out *os.File) error {
	defer out.Close()
	ciphertext := oracle.Encrypt(plaintext)

	_, err := out.Write(ciphertext)
	if err != nil {
		return err
	}

	return nil
}
