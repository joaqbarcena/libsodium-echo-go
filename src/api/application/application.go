package application

import (
	"github.com/gin-gonic/gin"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/usecase/encryption"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/usecase/response"
	handler "github.com/joaqbarcena/libsodium-echo-go/src/api/infrastructure/adapter/entrypoint/rest"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/infrastructure/adapter/provider"
)

func Run() {
	router := gin.Default()
	filesProvider := provider.NewFilesProvider()

	encryptionUseCase := encryption.NewEncryptionUseCase()
	fileResponseUseCase := response.NewFileResponseUseCase(filesProvider)
	fileResponseEncryptedUseCase := response.NewFileResponseEncryptedUseCase(filesProvider, encryptionUseCase)

	encryptionHandler := handler.NewEncryptionHandler(encryptionUseCase)
	nativeMockHandler := handler.NewNativeMockHandler(fileResponseUseCase, fileResponseEncryptedUseCase)

	handler.Wire(router,
		encryptionHandler.GetHandlerConfig(),
		nativeMockHandler.GetHandlerConfig(),
	)

	router.Run()
}
