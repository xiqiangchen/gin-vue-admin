package pricer

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"

	"github.com/rogpeppe/fastuuid"
)

const (
	ivSize        = 16
	textSize      = 8
	signatureSize = 4
	totalSize     = ivSize + textSize + signatureSize

	ivOffset        = 0
	textOffset      = ivOffset + ivSize
	signatureOffset = textOffset + textSize
)

var DefaultPricer, _ = New("dac78876796e5479fc66ca252790cb1c3cb2687a04f205acb54ccbdb8504b775",
	"a6b120682f56d2b7df82087b79bee9cd2435bbfa6f2dbf35638aeafc039350ee")

type Decrypter struct {
	*fastuuid.Generator
	EKey []byte // 加密 key
	IKey []byte // 校验 key
}

// base64URLEncode performs base64 URL encoding on byte slices.
func base64URLEncode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

// base64URLDecode decodes a base64 URL encoded string.
func base64URLDecode(src string) ([]byte, error) {
	// Correct padding for base64 URL encoding
	for len(src)%4 != 0 {
		src += "="
	}
	return base64.URLEncoding.DecodeString(src)
}

// hmacEncode generates HMAC with SHA1.
func hmacEncode(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	h.Write(data)
	return h.Sum(nil)
}

// New creates a new Decrypter with encryption and integrity keys.
func New(ekey, ikey string) (*Decrypter, error) {
	ek, err := hex.DecodeString(ekey)
	if err != nil {
		return nil, fmt.Errorf("invalid encryption key: %w", err)
	}
	ik, err := hex.DecodeString(ikey)
	if err != nil {
		return nil, fmt.Errorf("invalid integrity key: %w", err)
	}
	g, err := fastuuid.NewGenerator()
	if err != nil {
		return nil, fmt.Errorf("failed to create UUID generator: %w", err)
	}

	return &Decrypter{
		Generator: g,
		EKey:      ek,
		IKey:      ik,
	}, nil
}

// Decode decrypts the input string into an int64 value.
func (d *Decrypter) Decode(src string) (float64, error) {
	dest, err := base64URLDecode(src)
	if err != nil {
		return 0, fmt.Errorf("base64 decode error: %w", err)
	}
	if len(dest) != totalSize {
		return 0, errors.New("invalid input length")
	}

	iv := dest[ivOffset:textOffset]
	encText := dest[textOffset:signatureOffset]
	signature := dest[signatureOffset:]

	// Generate HMAC for decryption
	pad := hmacEncode(d.EKey, iv)
	plainText := make([]byte, textSize)
	for i := 0; i < textSize; i++ {
		plainText[i] = encText[i] ^ pad[i]
	}

	var text int64
	if err := binary.Read(bytes.NewReader(plainText), binary.BigEndian, &text); err != nil {
		return 0, fmt.Errorf("binary read error: %w", err)
	}

	// Verify signature
	bText := append(plainText, iv...)
	expectedSig := hmacEncode(d.IKey, bText)[:signatureSize]
	if !bytes.Equal(expectedSig, signature) {
		return 0, errors.New("signature verification failed")
	}
	return float64(text) / 100, nil
}

// Encode encrypts the int64 value into a base64 URL encoded string.
func (d *Decrypter) Encode(text float64) (string, error) {
	//iv := d.Next()[:ivSize]
	ivTmp := d.Next()
	iv := ivTmp[:ivSize]

	num := int64(math.Round(text * 100))

	// Convert int64 to bytes
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, num); err != nil {
		return "", fmt.Errorf("binary write error: %w", err)
	}
	plainText := buf.Bytes()

	// Encrypt using HMAC
	pad := hmacEncode(d.EKey, iv)
	encText := make([]byte, textSize)
	for i := 0; i < textSize; i++ {
		encText[i] = pad[i] ^ plainText[i]
	}

	// Generate signature
	bText := append(plainText, iv...)
	signature := hmacEncode(d.IKey, bText)[:signatureSize]

	// Concatenate iv, encrypted text, and signature
	result := append(append(iv, encText...), signature...)
	return base64URLEncode(result)[:38], nil
}
