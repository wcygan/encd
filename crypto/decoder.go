package crypto

import "io"

func Decode(ciphertext []byte, password string, out io.Writer) error {
	oracle, err := NewOracle(password)
	if err != nil {
		return err
	}

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
