package utils

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "errors"
    "io"
)

var ErrInvalidKey = errors.New("geçersiz şifreleme anahtarı")

// AES-256-GCM ile şifrele (KVKK hassas veri)
func Encrypt(plain string, hexKey string) (string, error) {
    key, err := hex.DecodeString(hexKey)
    if err != nil || len(key) != 32 {
        return "", ErrInvalidKey
    }
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }
    ct := gcm.Seal(nonce, nonce, []byte(plain), nil)
    return hex.EncodeToString(ct), nil
}

func Decrypt(cipherHex string, hexKey string) (string, error) {
    key, err := hex.DecodeString(hexKey)
    if err != nil || len(key) != 32 {
        return "", ErrInvalidKey
    }
    data, err := hex.DecodeString(cipherHex)
    if err != nil {
        return "", err
    }
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    nonceSize := gcm.NonceSize()
    if len(data) < nonceSize {
        return "", errors.New("şifreli veri çok kısa")
    }
    nonce, ct := data[:nonceSize], data[nonceSize:]
    plain, err := gcm.Open(nil, nonce, ct, nil)
    if err != nil {
        return "", err
    }
    return string(plain), nil
}
