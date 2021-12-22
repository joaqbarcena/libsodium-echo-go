package encryption

import (
	"encoding/base64"
	"fmt"

	"github.com/GoKillers/libsodium-go/cryptobox"
	"github.com/GoKillers/libsodium-go/sodium"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/model"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/resilience"
)

type UseCase interface {
	Execute(model.EncryptInfo) (string, *resilience.ApiError)
}

type implementation struct{}

func NewEncryptionUseCase() *implementation {
	sodium.Init()
	return &implementation{}
}

func (impl *implementation) Execute(encryptInfo model.EncryptInfo) (string, *resilience.ApiError) {
	decodedPk, decodeError := base64.StdEncoding.DecodeString(encryptInfo.EncodedPublicKey)

	if decodeError != nil {
		fmt.Printf("Couldnt decode base64: %s", decodeError.Error())
		return "", resilience.NewApiError(500, "cannot decode base64 pk")
	}

	encrypted, status := cryptobox.CryptoBoxSeal([]byte(encryptInfo.TextToEncrypt), decodedPk)
	if status == 0 {
		encodedEncryption := base64.StdEncoding.EncodeToString(encrypted)

		return encodedEncryption, nil
	} else {
		return "", resilience.NewApiError(500, fmt.Sprintf("cannot encrypt message, [exit status: %d]", status))
	}
}
