package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"golang.org/x/crypto/argon2"
)

const MinimumPasswordLength = 16
const salt = "eYPq4kRvyHYTdRykUsVbG+Azun77wwd2Th7yzjedHhc="

// The Oracle struct manages the encryption and decryption of data.
type Oracle struct {
	Aead  cipher.AEAD
	Nonce []byte
}

func (c *Oracle) Decrypt(ciphertext []byte) ([]byte, error) {
	return c.Aead.Open(nil, c.Nonce, ciphertext, nil)
}

func (c *Oracle) Encrypt(plaintext []byte) []byte {
	return c.Aead.Seal(nil, c.Nonce, plaintext, nil)
}

// NewOracle creates a new Oracle struct.
func NewOracle(passphrase string) (*Oracle, error) {
	o := new(Oracle)

	kdf := argon2.Key([]byte(passphrase), []byte(salt), 4, 32*1024, 4, 44)
	o.Nonce = kdf[32:]

	block, err := aes.NewCipher(kdf[:32])
	if err != nil {
		return o, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return o, err
	}

	o.Aead = aead

	return o, nil
}
