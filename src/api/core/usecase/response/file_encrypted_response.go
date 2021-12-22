package response

import (
	"regexp"
	"strings"

	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/model"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/resilience"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/usecase/encryption"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/infrastructure/adapter/provider"
)

type FileResponseEncryptedUseCase interface {
	Execute(model.FileResponseEncryptedInfo) (string, *resilience.ApiError)
}

type fileResponseEncryptedUseCaseImplementation struct {
	filesProvider     provider.FilesProvider
	encryptionUseCase encryption.UseCase
}

func NewFileResponseEncryptedUseCase(filesProvider provider.FilesProvider, encryptionUseCase encryption.UseCase) *fileResponseEncryptedUseCaseImplementation {
	return &fileResponseEncryptedUseCaseImplementation{
		filesProvider:     filesProvider,
		encryptionUseCase: encryptionUseCase,
	}
}

func (impl *fileResponseEncryptedUseCaseImplementation) Execute(info model.FileResponseEncryptedInfo) (string, *resilience.ApiError) {
	response, err := impl.filesProvider.GetFileAsString(fileAsResponse)

	if err != nil {
		return "", resilience.NewApiError(500, "Couldnt read file to get response")
	}

	re := regexp.MustCompile("<<(.*)>>")

	for {
		loc := re.FindStringIndex(response)
		if loc == nil {
			break
		}

		encryptInfo := model.EncryptInfo{
			EncodedPublicKey: info.EncodedPublicKey,
			TextToEncrypt:    response[loc[0]+2 : loc[1]-2],
		}
		encryptedText, encryptionErr := impl.encryptionUseCase.Execute(encryptInfo)

		if encryptionErr != nil {
			return response, encryptionErr
		}

		response = strings.Join([]string{
			response[:loc[0]],
			encryptedText,
			response[loc[1]:],
		}, "")
	}

	return response, nil
}
