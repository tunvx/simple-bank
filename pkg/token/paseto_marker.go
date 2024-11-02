package token

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/ed25519"
)

// PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto     *paseto.V2
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey
}

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker(publicKeyBase64 string, privateKeyBase64 string) (Maker, error) {
	// Decode public key
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("error decoding public key: %w", err)
	}
	publicKey := ed25519.PublicKey(publicKeyBytes)

	// Decode private key
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("error decoding private key: %w", err)
	}
	privateKey := ed25519.PrivateKey(privateKeyBytes)

	return &PasetoMaker{
		paseto:     paseto.NewV2(),
		publicKey:  publicKey,
		privateKey: privateKey,
	}, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(userID int64, role string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, role, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := maker.paseto.Sign(maker.privateKey, payload, nil)
	if err != nil {
		return "", nil, err
	}

	return token, payload, nil
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Verify(token, maker.publicKey, &payload, nil)
	if err != nil {
		return nil, err
	}

	if time.Now().After(payload.ExpiredAt) {
		return nil, errors.New("token has expired")
	}

	return payload, nil
}
