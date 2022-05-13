package crypto

import "io"

func Encode(plaintext []byte, password string, out io.Writer) error {
	oracle, err := NewOracle(password)
	if err != nil {
		return err
	}

	ciphertext := oracle.Encrypt(plaintext)

	_, err = out.Write(ciphertext)
	if err != nil {
		return err
	}

	return nil
}
