package response

import (
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/resilience"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/infrastructure/adapter/provider"
)

const (
	fileAsResponse = "response.json"
)

type FileResponseUseCase interface {
	Execute() (string, *resilience.ApiError)
}

type fileResponseUseCaseImplementation struct {
	filesProvider provider.FilesProvider
}

func NewFileResponseUseCase(filesProvider provider.FilesProvider) *fileResponseUseCaseImplementation {
	return &fileResponseUseCaseImplementation{
		filesProvider: filesProvider,
	}
}

func (impl *fileResponseUseCaseImplementation) Execute() (string, *resilience.ApiError) {
	response, err := impl.filesProvider.GetFileAsString(fileAsResponse)

	if err != nil {
		return "", resilience.NewApiError(500, "Couldnt read file to get response")
	}

	return response, nil
}
