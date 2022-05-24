package keypairgen

import (
	"encoding/base64"
	"fmt"

	"github.com/GoKillers/libsodium-go/cryptobox"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/model"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/resilience"
)

type UseCase interface {
	Execute(seed string) (*model.KeyPair, *resilience.ApiError)
}

type keyPairGenUseCase struct{}

func NewKeyPairGenUseCase() *keyPairGenUseCase {
	return &keyPairGenUseCase{}
}

func (keyPairGenUseCase *keyPairGenUseCase) Execute(seed string) (*model.KeyPair, *resilience.ApiError) {
	var secretKey, publicKey []byte
	var status int

	if len(seed) > 0 {
		decodedSeed, decodeError := base64.StdEncoding.DecodeString(seed)

		if decodeError != nil {
			fmt.Printf("Couldnt decode base64: %s", decodeError.Error())
			return nil, resilience.NewApiError(500, "cannot decode base64 seed")
		}

		secretKey, publicKey, status = cryptobox.CryptoBoxSeedKeyPair(decodedSeed)
	} else {
		secretKey, publicKey, status = cryptobox.CryptoBoxKeyPair()
	}

	if status != 0 {
		return nil, resilience.NewApiError(500, fmt.Sprintf("cannot generate key pair, [exit status: %d]", status))
	}

	encodedSecretKey := base64.StdEncoding.EncodeToString(secretKey)
	encodedPublicKey := base64.StdEncoding.EncodeToString(publicKey)

	return &model.KeyPair{
		SecretKey: encodedSecretKey,
		PublicKey: encodedPublicKey,
	}, nil
}
