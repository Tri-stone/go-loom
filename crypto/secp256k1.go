// +build evm

package crypto

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

type Secp256k1PrivateKey ecdsa.PrivateKey

func (s *Secp256k1PrivateKey) ToECDSAPrivKey() *ecdsa.PrivateKey {
	ecdsaPrivKey := ecdsa.PrivateKey(*s)
	return &ecdsaPrivKey
}

func LoadSecp256k1PrivKey(filePath string) (*Secp256k1PrivateKey, error) {
	privKey, err := crypto.LoadECDSA(filePath)
	if err != nil {
		return nil, err
	}

	secpPrivKey := Secp256k1PrivateKey(*privKey)

	return &secpPrivKey, nil
}
