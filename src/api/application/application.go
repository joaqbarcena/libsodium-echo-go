package application

import (
	"github.com/gin-gonic/gin"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/usecase/encryption"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/usecase/keypairgen"
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
	keyPairGenUseCase := keypairgen.NewKeyPairGenUseCase()

	encryptionHandler := handler.NewEncryptionHandler(encryptionUseCase)
	nativeMockHandler := handler.NewNativeMockHandler(fileResponseUseCase, fileResponseEncryptedUseCase)
	keyPairGenHandler := handler.NewKeyPairGenHandler(keyPairGenUseCase)

	handler.Wire(router,
		encryptionHandler.GetHandlerConfig(),
		nativeMockHandler.GetHandlerConfig(),
		keyPairGenHandler.GetHandlerConfig(),
	)

	router.Run()
}
