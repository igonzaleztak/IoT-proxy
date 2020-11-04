package cipherlib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"

	ecies "github.com/ecies/go"
)

// EncryptWithPublicKey encrypts a message using ECIES
func EncryptWithPublicKey(pubKeyECDSA ecdsa.PublicKey, msg []byte) ([]byte, error) {
	// Convert the public key to the appropiate format
	var pubKeyBytes []byte = elliptic.Marshal(pubKeyECDSA.Curve, pubKeyECDSA.X, pubKeyECDSA.Y)
	pubKey, err := ecies.NewPublicKeyFromBytes(pubKeyBytes)
	if err != nil {
		return nil, err
	}

	// Encrypt the message with the public key
	cipherText, err := ecies.Encrypt(pubKey, msg)
	if err != nil {
		return nil, err
	}

	return cipherText, nil
}

// DecryptWithPrivateKey Decrypts a text encrypted using ECIES
func DecryptWithPrivateKey(pkBytes *ecdsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	// Convert the private key to the format of the libary
	var privKey = ecies.NewPrivateKeyFromBytes(pkBytes.D.Bytes())

	// Decrypt the cipherText
	plainText, err := ecies.Decrypt(privKey, ciphertext)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// VerifySignature verifies the signature of a message
func VerifySignature(pubKey, msg, signature []byte) bool {
	return secp256k1.VerifySignature(pubKey, msg, signature)
}

//randBytes creates a lenght byte long array of bytes
func randBytes(length int) []byte {
	b := make([]byte, length)
	rand.Read(b)
	return b
}

// SymmetricEncryption ecnrypts some data with AES-256-gcm
func SymmetricEncryption(key, data []byte) ([]byte, error) {
	// Generate a new aes cipher using the 32 byte long key
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Generete the GCM cipher
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Generate a random nonce that will be used in the cipher
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// Cipher the data
	ciphertext := gcm.Seal(nil, nonce, data, nil)

	return append(nonce, ciphertext...), nil
}

// DecryptSymmetricEncryption decrypts aes-256-gcm
func DecryptSymmetricEncryption(key, data []byte) error {
	nonce := data[:12]
	ct := data[12:]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plainText, err := aesgcm.Open(nil, nonce, ct, nil)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(plainText))

	return nil
}

// HashData hashes data
func HashData(msg []byte) []byte {
	hash := sha256.Sum256([]byte(msg))
	return hash[:]
}

// SignData Sign with private key
func SignData(privateKey *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	signedData, err := secp256k1.Sign(HashData(data), crypto.FromECDSA(privateKey))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return signedData, nil
}
