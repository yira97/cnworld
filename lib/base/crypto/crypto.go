package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
const nonceLength = 12

// key 必须是 16, 24, or 32 字节中的一种
func newCipher(key []byte) cipher.Block {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	return block
}

func newGCM(key []byte, nonceLength int) cipher.AEAD {
	block := newCipher(key)

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	aesgcm, err := cipher.NewGCMWithNonceSize(block, nonceLength)
	if err != nil {
		panic(err.Error())
	}
	return aesgcm
}

// https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
func NewGCM_encrypt(key16 string, plainText []byte) []byte {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, _ := hex.DecodeString(key16)

	// populates our nonce with a cryptographically secure
	// random sequence
	nonce := make([]byte, nonceLength)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm := newGCM(key, nonceLength)

	// 第一个参数的意思是说, 生成的ciphertext 直接拼到 nonce后面, 一块返回
	// 真正decrypt, aesgcm.Open(nil, nonce, ciphertext, nil)  , 之前实际上已经把nonce剔掉了.
	ciphertext := aesgcm.Seal(nonce, nonce, plainText, nil)
	return ciphertext
}

// https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
func NewGCM_decrypt(key16 string, cipherText []byte) ([]byte, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, _ := hex.DecodeString(key16)

	aesgcm := newGCM(key, nonceLength)

	nonce, ciphertext := cipherText[:nonceLength], cipherText[nonceLength:]
	plainText, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
